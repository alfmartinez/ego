package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/render"
	"fmt"
	"image"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	RegisterRendererFactory("glfw", func(config configuration.Renderer) Renderer {
		display := display.CreateDisplay(config.Display)
		return &glfwRenderer{display: display}
	})
}

type glfwRenderer struct {
	display display.Display
	window  *glfw.Window
	prog    uint32
}

var (
	quad = []float32{
		-1, -1, 0,
		1, -1, 0,
		1, 1, 0,
		-1, 1, 0,
	}
)

// Init implements Renderer
func (g *glfwRenderer) Init() {
	runtime.LockOSThread()
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
	g.window.MakeContextCurrent()
	g.prog = initOpenGL()
}

// Refresh implements Renderer
func (g *glfwRenderer) Refresh() {
	img := g.display.Render()
	draw(img, g.prog)
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

func draw(img image.Image, prog uint32) {
	texture, err := makeTexture(img)
	if err != nil {
		panic(err)
	}
	vao := makeVao(quad)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)
	gl.BindVertexArray(vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(quad)/3))
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func makeTexture(img image.Image) (uint32, error) {

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}
