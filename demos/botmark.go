package main

import (
	"fmt"
	"github.com/ajhager/engi"
	"math/rand"
)

var (
	bots   []*Bot
	on     bool
	num    int
	region *engi.Region
	batch  *engi.Batch
	font   *engi.Font
)

type Bot struct {
	*engi.Sprite
	DX, DY float32
}

type Game struct {
	*engi.Game
}

func (game *Game) Preload() {
	engi.Files.Add("bot", "data/icon.png")
	engi.Files.Add("font", "data/font.png")
}

func (game *Game) Setup() {
	engi.SetBg(0x362d38)
	font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
	texture := engi.Files.Image("bot")
	region = engi.NewRegion(texture, 0, 0, int(texture.Width()), int(texture.Height()))
	batch = engi.NewBatch(engi.Width(), engi.Height())
}

func (game *Game) Update(dt float32) {
	if on {
		for i := 0; i < 10; i++ {
			bot := &Bot{engi.NewSprite(region, 0, 0), rand.Float32() * 500, rand.Float32()*500 - 250}
			bots = append(bots, bot)
		}
		num += 10
	}

	width := engi.Width()
	height := engi.Height()

	for _, bot := range bots {
		bot.Position.X += bot.DX * dt
		bot.Position.Y += bot.DY * dt
		bot.DY += 750 * dt

		if bot.Position.X < 0 {
			bot.DX *= -1
			bot.Position.X = 0
		} else if bot.Position.X > width {
			bot.DX *= -1
			bot.Position.X = width
		}

		if bot.Position.Y < 0 {
			bot.DY = 0
			bot.Position.Y = 0
		} else if bot.Position.Y > height {
			bot.DY *= -.85
			bot.Position.Y = height
			if rand.Float32() > 0.5 {
				bot.DY -= rand.Float32() * 200
			}
		}
	}
}

func (game *Game) Render() {
	batch.Begin()
	for _, bot := range bots {
		bot.Render(batch)
	}

	sfps := fmt.Sprintf("%vFPS", int(engi.Time.Fps()))
	snum := fmt.Sprintf("#%v", int(num))
	font.Print(batch, sfps, 0, 0, 0xffffff)
	font.Print(batch, snum, 120, 0, 0xffffff)

	batch.End()
}

func (game *Game) Mouse(x, y float32, action engi.Action) {
	switch action {
	case engi.PRESS:
		on = true
	case engi.RELEASE:
		on = false
	}
}

func (game *Game) Resize(w, h float32) {
	batch.SetProjection(w, h)
}

func main() {
	engi.Open("Botmark", 800, 600, false, &Game{})
}
