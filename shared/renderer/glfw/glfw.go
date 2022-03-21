package glfw

import (
	"ego/engine/context"
	"ego/engine/display"
	"ego/engine/input"
	"ego/engine/render"
	"ego/engine/renderer"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/spf13/viper"
)

func Register() {
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
func (g *glfwRenderer) Start() {
	handler := context.GetContext().Get("input").(input.InputHandler)
	g.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		handler.Handle(input.Event{
			Key:    ConvertKey(key),
			Action: ConvertAction(action),
		})
	})
}

// Render implements Renderer
func (g *glfwRenderer) Render(i interface{}) {
	c := i.(render.Convertible)
	s := render.ConvertObjectToDisplayable(c)
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
