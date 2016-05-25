package main

import (
	"github.com/ajhager/engi"
)

type Game struct {
	*engi.Game
	bot   engi.Drawable
	batch *engi.Batch
	font  *engi.Font
}

func (game *Game) Preload() {
	engi.Files.Add("bot", "data/icon.png")
	engi.Files.Add("font", "data/font.png")
}

func (game *Game) Setup() {
	engi.SetBg(0x362d38)
	game.bot = engi.Files.Image("bot")
	game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
	game.batch = engi.NewBatch(engi.Width(), engi.Height())
}

func (game *Game) Render() {
	game.batch.Begin()
	midx := engi.Width() / 2
	midy := engi.Height() / 2
	game.font.Print(game.batch, "ENGI", midx-37, midy-120, 0xffffff)
	game.batch.Draw(game.bot, midx, midy, 0.5, 0.5, 10, 10, 0, 0xffffff, 1)
	game.batch.End()
}

func (game *Game) Resize(w, h float32) {
	game.batch.SetProjection(w, h)
}

func main() {
	engi.Open("Hello", 800, 600, false, &Game{})
}
