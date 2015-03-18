// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engi

type Config struct {
	Title      string
	Width      int
	Height     int
	Fullscreen bool
}

func DefaultConfig() Config {
	return Config{
		Title:      "ENGi",
		Width:      1024,
		Height:     640,
		Fullscreen: false,
	}
}

type Application interface {
	Config() Config
	Create()
	Resize(width, height int)
	Render()
	Pause()
	Resume()
	Dispose()
}

type App struct{}

func (app *App) Config() Config {
	return DefaultConfig()
}

func (app *App) Create()                  {}
func (app *App) Resize(width, height int) {}
func (app *App) Render()                  {}
func (app *App) Pause()                   {}
func (app *App) Resume()                  {}
func (app *App) Dispose()                 {}
