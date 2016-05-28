// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !netgo,!android

package engi

import (
	"image"
	"image/draw"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/ajhager/webgl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var window *glfw.Window

// fatalErr calls log.Fatal with the given error if it is non-nil.
func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	windowWidth  int
	windowHeight int
)

func init() {
	runtime.LockOSThread()
}

func run(title string, width, height int, fullscreen bool) {
	fatalErr(glfw.Init())

	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	if fullscreen {
		width = mode.Width
		height = mode.Height
		glfw.WindowHint(glfw.Decorated, glfw.False)
	} else {
		monitor = nil
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, title, monitor, nil)
	fatalErr(err)
	window.MakeContextCurrent()

	if !fullscreen {
		window.SetPos((mode.Width-width)/2, (mode.Height-height)/2)
	}

	windowWidth, windowHeight = window.GetSize()
	width, height = window.GetFramebufferSize()

	glfw.SwapInterval(1)

	gl = webgl.NewContext()

	gl.Viewport(0, 0, width, height)
	window.SetFramebufferSizeCallback(func(window *glfw.Window, w, h int) {
		windowWidth, windowHeight = window.GetSize()
		gl.Viewport(0, 0, w, h)
		responder.Resize(float32(windowWidth), float32(windowHeight))
	})

	window.SetCursorPosCallback(func(window *glfw.Window, x, y float64) {
		responder.Mouse(float32(x), float32(y), MOVE)
	})

	window.SetMouseButtonCallback(func(window *glfw.Window, b glfw.MouseButton, a glfw.Action, m glfw.ModifierKey) {
		x, y := window.GetCursorPos()
		if a == glfw.Press {
			responder.Mouse(float32(x), float32(y), PRESS)
		} else {
			responder.Mouse(float32(x), float32(y), RELEASE)
		}
	})

	window.SetScrollCallback(func(window *glfw.Window, xoff, yoff float64) {
		responder.Scroll(float32(yoff))
	})

	window.SetKeyCallback(func(window *glfw.Window, k glfw.Key, s int, a glfw.Action, m glfw.ModifierKey) {
		switch a {
		case glfw.Press:
			responder.Key(Key(k), Modifier(m), PRESS)
		case glfw.Release:
			responder.Key(Key(k), Modifier(m), RELEASE)
		case glfw.Repeat:
			responder.Key(Key(k), Modifier(m), REPEAT)
		}
	})

	window.SetCharCallback(func(window *glfw.Window, char rune) {
		responder.Type(char)
	})

	responder.Preload()
	Files.Load(func() {})
	responder.Setup()

	shouldClose := window.ShouldClose()
	for !shouldClose {
		responder.Update(Time.Delta())
		gl.Clear(gl.COLOR_BUFFER_BIT)
		responder.Render()
		window.SwapBuffers()
		glfw.PollEvents()
		Time.Tick()

		shouldClose = window.ShouldClose()
	}
	window.Destroy()
	glfw.Terminate()
	responder.Close()
}

func width() float32 {
	return float32(windowWidth)
}

func height() float32 {
	return float32(windowHeight)
}

func exit() {
	window.SetShouldClose(true)
}

func init() {
	Dash = Key(glfw.KeyMinus)
	Apostrophe = Key(glfw.KeyApostrophe)
	Semicolon = Key(glfw.KeySemicolon)
	Equals = Key(glfw.KeyEqual)
	Comma = Key(glfw.KeyComma)
	Period = Key(glfw.KeyPeriod)
	Slash = Key(glfw.KeySlash)
	Backslash = Key(glfw.KeyBackslash)
	Backspace = Key(glfw.KeyBackspace)
	Tab = Key(glfw.KeyTab)
	CapsLock = Key(glfw.KeyCapsLock)
	Space = Key(glfw.KeySpace)
	Enter = Key(glfw.KeyEnter)
	Escape = Key(glfw.KeyEscape)
	Insert = Key(glfw.KeyInsert)
	PrintScreen = Key(glfw.KeyPrintScreen)
	Delete = Key(glfw.KeyDelete)
	PageUp = Key(glfw.KeyPageUp)
	PageDown = Key(glfw.KeyPageDown)
	Home = Key(glfw.KeyHome)
	End = Key(glfw.KeyEnd)
	Pause = Key(glfw.KeyPause)
	ScrollLock = Key(glfw.KeyScrollLock)
	ArrowLeft = Key(glfw.KeyLeft)
	ArrowRight = Key(glfw.KeyRight)
	ArrowDown = Key(glfw.KeyDown)
	ArrowUp = Key(glfw.KeyUp)
	LeftBracket = Key(glfw.KeyLeftBracket)
	LeftShift = Key(glfw.KeyLeftShift)
	LeftControl = Key(glfw.KeyLeftControl)
	LeftSuper = Key(glfw.KeyLeftSuper)
	LeftAlt = Key(glfw.KeyLeftAlt)
	RightBracket = Key(glfw.KeyRightBracket)
	RightShift = Key(glfw.KeyRightShift)
	RightControl = Key(glfw.KeyRightControl)
	RightSuper = Key(glfw.KeyRightSuper)
	RightAlt = Key(glfw.KeyRightAlt)
	Zero = Key(glfw.Key0)
	One = Key(glfw.Key1)
	Two = Key(glfw.Key2)
	Three = Key(glfw.Key3)
	Four = Key(glfw.Key4)
	Five = Key(glfw.Key5)
	Six = Key(glfw.Key6)
	Seven = Key(glfw.Key7)
	Eight = Key(glfw.Key8)
	Nine = Key(glfw.Key9)
	F1 = Key(glfw.KeyF1)
	F2 = Key(glfw.KeyF2)
	F3 = Key(glfw.KeyF3)
	F4 = Key(glfw.KeyF4)
	F5 = Key(glfw.KeyF5)
	F6 = Key(glfw.KeyF6)
	F7 = Key(glfw.KeyF7)
	F8 = Key(glfw.KeyF8)
	F9 = Key(glfw.KeyF9)
	F10 = Key(glfw.KeyF10)
	F11 = Key(glfw.KeyF11)
	F12 = Key(glfw.KeyF12)
	A = Key(glfw.KeyA)
	B = Key(glfw.KeyB)
	C = Key(glfw.KeyC)
	D = Key(glfw.KeyD)
	E = Key(glfw.KeyE)
	F = Key(glfw.KeyF)
	G = Key(glfw.KeyG)
	H = Key(glfw.KeyH)
	I = Key(glfw.KeyI)
	J = Key(glfw.KeyJ)
	K = Key(glfw.KeyK)
	L = Key(glfw.KeyL)
	M = Key(glfw.KeyM)
	N = Key(glfw.KeyN)
	O = Key(glfw.KeyO)
	P = Key(glfw.KeyP)
	Q = Key(glfw.KeyQ)
	R = Key(glfw.KeyR)
	S = Key(glfw.KeyS)
	T = Key(glfw.KeyT)
	U = Key(glfw.KeyU)
	V = Key(glfw.KeyV)
	W = Key(glfw.KeyW)
	X = Key(glfw.KeyX)
	Y = Key(glfw.KeyY)
	Z = Key(glfw.KeyZ)
	NumLock = Key(glfw.KeyNumLock)
	NumMultiply = Key(glfw.KeyKPMultiply)
	NumDivide = Key(glfw.KeyKPDivide)
	NumAdd = Key(glfw.KeyKPAdd)
	NumSubtract = Key(glfw.KeyKPSubtract)
	NumZero = Key(glfw.KeyKP0)
	NumOne = Key(glfw.KeyKP1)
	NumTwo = Key(glfw.KeyKP2)
	NumThree = Key(glfw.KeyKP3)
	NumFour = Key(glfw.KeyKP4)
	NumFive = Key(glfw.KeyKP5)
	NumSix = Key(glfw.KeyKP6)
	NumSeven = Key(glfw.KeyKP7)
	NumEight = Key(glfw.KeyKP8)
	NumNine = Key(glfw.KeyKP9)
	NumDecimal = Key(glfw.KeyKPDecimal)
	NumEnter = Key(glfw.KeyKPEnter)
}

func NewImageObject(img *image.NRGBA) *ImageObject {
	return &ImageObject{img}
}

type ImageObject struct {
	data *image.NRGBA
}

func (i *ImageObject) Data() interface{} {
	return i.data
}

func (i *ImageObject) Width() int {
	return i.data.Rect.Max.X
}

func (i *ImageObject) Height() int {
	return i.data.Rect.Max.Y
}

func loadImage(r Resource) (Image, error) {
	file, err := os.Open(r.url)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	b := img.Bounds()
	newm := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(newm, newm.Bounds(), img, b.Min, draw.Src)

	return &ImageObject{newm}, nil
}

func loadJson(r Resource) (string, error) {
	file, err := ioutil.ReadFile(r.url)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

type Assets struct {
	queue  []string
	cache  map[string]Image
	loads  int
	errors int
}

func NewAssets() *Assets {
	return &Assets{make([]string, 0), make(map[string]Image), 0, 0}
}

func (a *Assets) Image(path string) {
	a.queue = append(a.queue, path)
}

func (a *Assets) Get(path string) Image {
	return a.cache[path]
}

func (a *Assets) Load(onFinish func()) {
	if len(a.queue) == 0 {
		onFinish()
	} else {
		for _, path := range a.queue {
			img := LoadImage(path)
			a.cache[path] = img
		}
	}
}

func LoadImage(data interface{}) Image {
	var m image.Image

	switch data := data.(type) {
	default:
		log.Fatal("NewTexture needs a string or io.Reader")
	case string:
		file, err := os.Open(data)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		m = img
	case io.Reader:
		img, _, err := image.Decode(data)
		if err != nil {
			log.Fatal(err)
		}
		m = img
	case image.Image:
		m = data
	}

	b := m.Bounds()
	newm := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(newm, newm.Bounds(), m, b.Min, draw.Src)

	return &ImageObject{newm}
}
