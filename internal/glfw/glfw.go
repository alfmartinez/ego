package ego_glfw

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/render"
	"ego/pkg/renderer"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	renderer.RegisterRendererFactory("glfw", func(config configuration.Renderer) renderer.Renderer {
		display := display.CreateDisplay(config.Display)
		gl := CreateGLEngine()
		return &glfwRenderer{display: display, gl: gl}
	})
}

type glfwRenderer struct {
	display display.Display
	gl      GLEngine
	window  *glfw.Window
	exit    chan bool
}

// Init implements Renderer
func (g *glfwRenderer) Init() {
	g.display.Init()
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	//glfw.WindowHint(glfw.AutoIconify, glfw.True)

	window, err := glfw.CreateWindow(1920, 1080, "Ego", glfw.GetPrimaryMonitor(), nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	g.gl.Init()
	g.window = window
}

// Close implements Renderer
func (g *glfwRenderer) Close() {
	defer glfw.Terminate()
}

// IsAsync implements Renderer
func (*glfwRenderer) IsAsync() bool {
	return true
}

// Start implements Renderer
func (g *glfwRenderer) Start(exit chan bool) {
	g.exit = exit
	g.window.SetKeyCallback(g.GetKeyCallBack())
}

// Render implements Renderer
func (g *glfwRenderer) Render(tree render.RenderTree) {
	tree.Apply(func(node render.RenderNode) {
		s := node.Display()
		g.display.AddObject(s)
	})
}

// Refresh implements Renderer
func (g *glfwRenderer) Refresh() {
	img := g.display.Render()
	texture := g.gl.AddTexture(img)
	g.gl.Draw(texture)
	g.window.SwapBuffers()
	glfw.PollEvents()
}

func (g *glfwRenderer) GetKeyCallBack() glfw.KeyCallback {
	return func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape && action == glfw.Press {
			go func() {
				g.exit <- true
			}()
		}
	}
}
