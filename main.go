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

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

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

	verts := []float32{
		-0.5, -0.5,
		0.0, 0.5,
		0.5, -0.5}

	//Test gl
	var vBuff uint32
	gl.GenBuffers(1, &vBuff)
	gl.BindBuffer(gl.ARRAY_BUFFER, vBuff)

	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	gl.ClearColor(0, 0, 0, 1)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
