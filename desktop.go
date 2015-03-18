// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js,!android

package engi

import (
	"log"
	"runtime"

	"github.com/ajhager/engi/graphics"
	"github.com/ajhager/engi/graphics/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var window *glfw.Window

func Init(app Application) {
	config := app.Config()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}

	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	width := config.Width
	height := config.Height

	if config.Fullscreen {
		width = mode.Width
		height = mode.Height
		glfw.WindowHint(glfw.Decorated, 0)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, config.Title, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()

	if !config.Fullscreen {
		window.SetPos((mode.Width-width)/2, (mode.Height-height)/2)
	}

	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	width, height = window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	graphics.Init()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		app.Render()

		window.SwapBuffers()
		glfw.PollEvents()

		graphics.Tick()
	}

	window.Destroy()
	glfw.Terminate()
}
