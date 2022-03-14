package main // import "github.com/go-gl/example/gl41core-cube"

import (
	"image"
	_ "image/png"
	"log"
	"runtime"

	ego_glfw "ego/internal/glfw"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glEngine := ego_glfw.CreateGLEngine()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	glEngine.Init()

	for {
		texture := glEngine.AddTexture(image.NewNRGBA(image.Rect(0, 0, 800, 600)))
		glEngine.Draw(texture)

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
	defer glfw.Terminate()
}
