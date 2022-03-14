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

	window, err := glfw.CreateWindow(800, 600, "Ego", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	g.gl.Init()
	g.window = window
}

// Close implements Renderer
func (*glfwRenderer) Close() {
	defer glfw.Terminate()
}

// IsAsync implements Renderer
func (*glfwRenderer) IsAsync() bool {
	return true
}

// Start implements Renderer
func (g *glfwRenderer) Start(chan bool) {

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
