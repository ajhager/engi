// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/ajhager/engi"
	"github.com/ajhager/engi/graphics"
)

type Hello struct {
	*engi.App
}

func (app *Hello) Render() {
	println(graphics.Delta())
}

func main() {
	engi.Init(new(Hello))
}
