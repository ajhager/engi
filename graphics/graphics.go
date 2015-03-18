// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphics

import (
	"math"
	"time"
)

var (
	fps   float64
	delta float64

	elapsed float64
	frames  uint64
	start   time.Time
	frame   time.Time
)

func Init() {
	start = time.Now()
	Tick()
}

func Tick() {
	now := time.Now()
	frames += 1
	delta = now.Sub(frame).Seconds()
	elapsed += delta
	frame = now

	if elapsed >= 1 {
		fps = float64(frames)
		elapsed = math.Mod(elapsed, 1)
		frames = 0
	}
}

func Delta() float32 {
	return float32(delta)
}

func Fps() float32 {
	return float32(fps)
}

func Time() float32 {
	return float32(time.Now().Sub(start).Seconds())
}
