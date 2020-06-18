package main

import (
	"runtime"

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
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
