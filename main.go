package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, winErr := glfw.CreateWindow(1270, 720, "Hello World", nil, nil)
	if winErr != nil {
		panic(winErr)
	}
	window.MakeContextCurrent()

	//Init opengl
	glErr := gl.Init()
	if glErr != nil {
		panic(glErr)
	}

	gl.ClearColor(1, 1, 1, 1)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
