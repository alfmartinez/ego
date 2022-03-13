package glfw

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/render"
	"ego/pkg/renderer"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	renderer.RegisterRendererFactory("glfw", func(config configuration.Renderer) renderer.Renderer {
		display := display.CreateDisplay(config.Display)
		gl := CreateGlRenderer()
		return &glfwRenderer{display: display, gl: gl}
	})
}

type glfwRenderer struct {
	display display.Display
	gl      GlRenderer
	window  *glfw.Window
}

// Init implements Renderer
func (g *glfwRenderer) Init() {
	runtime.LockOSThread()
	g.display.Init()
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	g.window = window
}

// Close implements Renderer
func (*glfwRenderer) Close() {
	defer glfw.Terminate()
}

// IsAsync implements Renderer
func (*glfwRenderer) IsAsync() bool {
	return false
}

// Start implements Renderer
func (g *glfwRenderer) Start(chan bool) {
	g.gl.InitGl()
	g.window.MakeContextCurrent()

}

// Refresh implements Renderer
func (g *glfwRenderer) Refresh() {
	img := g.display.Render()
	g.gl.Draw(img)
	g.window.SwapBuffers()
	glfw.PollEvents()
}

// Render implements Renderer
func (g *glfwRenderer) Render(tree render.RenderTree) {
	tree.Apply(func(node render.RenderNode) {
		s := node.Display()
		g.display.AddObject(s)
	})
}
