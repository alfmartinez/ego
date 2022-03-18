package glfw

import (
	"ego/pkg/context"
	"ego/pkg/display"
	"ego/pkg/render"
	"ego/pkg/renderer"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/spf13/viper"
)

func init() {
	renderer.RegisterRendererFactory("glfw", func() renderer.Renderer {
		display := display.CreateDisplay()
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
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	window, err := glfw.CreateWindow(
		viper.GetInt("renderer.display.viewport.x"),
		viper.GetInt("renderer.display.viewport.y"),
		"Ego", glfw.GetPrimaryMonitor(), nil,
	)
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
func (g *glfwRenderer) Render(i interface{}) {
	s := render.ConvertObjectToDisplayable(i)
	g.display.AddObject(s)
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
