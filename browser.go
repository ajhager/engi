// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js

package engi

import (
	"log"
	"strconv"

	"github.com/ajhager/engi/graphics"
	"github.com/ajhager/engi/graphics/gl"
	"github.com/gopherjs/gopherjs/js"
)

var raf = js.Global.Get("requestAnimationFrame")

var canvas *js.Object
var app Application

func Init(application Application) {
	app = application
	config := app.Config()

	document := js.Global.Get("document")
	canvas = document.Call("createElement", "canvas")

	target := document.Call("getElementById", config.Title)
	if target == nil {
		target = document.Get("body")
	}
	target.Call("appendChild", canvas)

	glAttrs := map[string]bool{
		"alpha":                 false,
		"depth":                 false,
		"premultipliedAlpha":    false,
		"preserveDrawingBuffer": false,
		"antialias":             false,
	}

	context := canvas.Call("getContext", "webgl", glAttrs)
	if err := gl.Init(context); err != nil {
		log.Println(err)
	}

	var toPx = func(n int) string {
		return strconv.FormatInt(int64(n), 10) + "px"
	}

	canvas.Get("style").Set("display", "block")

	winWidth := js.Global.Get("innerWidth").Int()
	winHeight := js.Global.Get("innerHeight").Int()
	width := config.Width
	height := config.Height

	if config.Fullscreen {
		width = winWidth
		height = winHeight
	} else {
		canvas.Get("style").Set("marginLeft", toPx((winWidth-config.Width)/2))
		canvas.Get("style").Set("marginTop", toPx((winHeight-config.Height)/2))
	}

	pixelRatio := 1.0
	if js.Global.Get("devicePixelRatio") != js.Undefined {
		pixelRatio = js.Global.Get("devicePixelRatio").Float()
	}

	clientWidth := int32(float64(width) * pixelRatio)
	clientHeight := int32(float64(height) * pixelRatio)

	canvas.Get("style").Set("width", toPx(width))
	canvas.Get("style").Set("height", toPx(height))
	canvas.Set("width", clientWidth)
	canvas.Set("height", clientHeight)

	gl.Viewport(0, 0, clientWidth, clientHeight)
	graphics.Init()
	raf.Invoke(loop)
}

func loop() {
	raf.Invoke(loop)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	app.Render()
	graphics.Tick()
}
