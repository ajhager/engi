package main

import (
	"fmt"
	"github.com/ajhager/engi"
	"math"
)

type Entity struct {
	*engi.Sprite
	Velocity *engi.Point
	LastPos  *engi.Point
}

func NewEntity(region *engi.Region) *Entity {
	entity := &Entity{
		engi.NewSprite(region, 0, 0),
		&engi.Point{0, 0},
		&engi.Point{0, 0},
	}
	entity.Anchor.X = 0.5
	entity.Anchor.Y = 0.5
	return entity
}

func (e *Entity) Bounds() (float32, float32, float32, float32) {
	hWidth := e.Width() / 2
	hHeight := e.Height() / 2
	right := e.Position.X + hWidth
	left := e.Position.X - hWidth
	top := e.Position.Y - hHeight
	bottom := e.Position.Y + hHeight
	return right, left, top, bottom
}

func (a *Entity) Intersects(b *Entity) bool {
	aRight, aLeft, aTop, aBottom := a.Bounds()
	bRight, bLeft, bTop, bBottom := b.Bounds()
	return !(bLeft > aRight || aLeft > bRight || bBottom < aTop || bTop > aBottom)
}

var (
	ball      *Entity
	paddle    *Entity
	bricks    []*Entity
	font      *engi.Font
	lives     int
	livesText string
	score     int
	scoreText string
	speed     float32
	launching bool
	batch     *engi.Batch
)

type Game struct {
	*engi.Game
}

func (g *Game) Preload() {
	engi.Files.Add("ball", "data/ball.png")
	engi.Files.Add("bricks", "data/bricks.png")
	engi.Files.Add("paddle", "data/paddle.png")
	engi.Files.Add("font", "data/font.png")
}

func (g *Game) Setup() {
	batch = engi.NewBatch(engi.Width(), engi.Height())

	engi.SetBg(0x362d38)

	texture := engi.Files.Image("paddle")
	region := engi.NewRegion(texture, 0, 0, int(texture.Width()), int(texture.Height()))
	paddle = NewEntity(region)
	paddle.Position.X = engi.Width() / 2
	paddle.Position.Y = engi.Height() - 50

	texture = engi.Files.Image("ball")
	region = engi.NewRegion(texture, 0, 0, int(texture.Width()), int(texture.Height()))
	ball = NewEntity(region)
	ball.Position.X = paddle.Position.X
	ball.Position.Y = paddle.Position.Y - ball.Height() - 1

	font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)

	g.Reset()
}

func (g *Game) Resize(w, h float32) {
	batch.SetProjection(w, h)
}

func (g *Game) Key(key engi.Key, modifier engi.Modifier, action engi.Action) {
	if action == engi.PRESS {
		switch key {
		case engi.ArrowLeft:
			paddle.Velocity.X -= speed
		case engi.ArrowRight:
			paddle.Velocity.X += speed
		case engi.Space:
			if launching {
				launching = false
				ball.Velocity.Y = -speed
			}
		}
	} else if action == engi.RELEASE {
		switch key {
		case engi.ArrowLeft:
			paddle.Velocity.X += speed
		case engi.ArrowRight:
			paddle.Velocity.X -= speed
		}
	}
}

func (g *Game) Update(dt float32) {
	// Move paddle
	paddle.Position.X += paddle.Velocity.X * dt

	// Paddle vs World
	hPaddleWidth := paddle.Width() / 2

	if paddle.Position.X > engi.Width()-hPaddleWidth {
		paddle.Position.X = engi.Width() - hPaddleWidth
	}

	if paddle.Position.X < hPaddleWidth {
		paddle.Position.X = hPaddleWidth
	}

	// Sync launching ball
	if launching {
		ball.Position.X = paddle.Position.X
		return
	}

	// Move ball
	ball.LastPos.X = ball.Position.X
	ball.LastPos.Y = ball.Position.Y
	ball.Position.X += ball.Velocity.X * dt
	ball.Position.Y += ball.Velocity.Y * dt

	// Ball vs World
	hBallWidth := ball.Width() / 2

	if ball.Position.X > engi.Width()-hBallWidth {
		ball.Position.X = engi.Width() - hBallWidth
		ball.Velocity.X *= -1
	}

	if ball.Position.X < hBallWidth {
		ball.Position.X = hBallWidth
		ball.Velocity.X *= -1
	}

	hBallHeight := ball.Height() / 2

	if ball.Position.Y > engi.Height()-hBallHeight {
		g.BallLost()
	}

	if ball.Position.Y < hBallHeight {
		ball.Position.Y = hBallHeight
		ball.Velocity.Y *= -1
	}

	// Ball vs Paddle
	ballXMin := ball.Position.X - hBallWidth
	ballXMax := ball.Position.X + hBallWidth
	ballYMax := ball.Position.Y + hBallHeight
	paddleXMin := paddle.Position.X - hPaddleWidth
	paddleXMax := paddle.Position.X + hPaddleWidth
	paddleYMin := paddle.Position.Y - paddle.Height()/2

	if ballYMax > paddleYMin {
		if (ballXMax > paddleXMin) && (ballXMin < paddleXMax) {
			ball.Position.Y = paddleYMin - hBallHeight - 1

			intersect := (ball.Position.X - paddle.Position.X) / hPaddleWidth
			bounceAngle := float64(intersect*1.047 - 1.571)

			ball.Velocity.X = float32(float64(speed) * math.Cos(bounceAngle))
			ball.Velocity.Y = float32(float64(speed) * math.Sin(bounceAngle))
		}
	}

	// Ball vs Brick
	newBricks := make([]*Entity, 0)
	for _, brick := range bricks {
		if !ball.Intersects(brick) {
			newBricks = append(newBricks, brick)
		} else {
			g.BallHitBrick(ball, brick)
		}
	}
	bricks = newBricks
}

func (g *Game) Render() {
	batch.Begin()
	paddle.Render(batch)
	ball.Render(batch)

	for _, brick := range bricks {
		brick.Render(batch)
	}

	font.Print(batch, scoreText, 20, 570, 0xffffff)
	font.Print(batch, livesText, 620, 570, 0xffffff)

	batch.End()
}

func (g *Game) BallLost() {
	launching = true

	ball.Velocity.X = 0
	ball.Velocity.Y = 0
	ball.Position.X = paddle.Position.X
	ball.Position.Y = paddle.Position.Y - ball.Height() - 1

	lives -= 1
	livesText = fmt.Sprintf("lives: %d", lives)
	if lives == 0 {
		g.Reset()
	}
}

func (g *Game) Reset() {
	launching = true
	speed = 350
	lives = 3
	livesText = fmt.Sprintf("lives: %d", lives)
	score = 0
	scoreText = fmt.Sprintf("score: %d", score)

	texture := engi.Files.Image("bricks")
	width := int(texture.Width())
	height := int(texture.Height()) / 4
	brickRegions := []*engi.Region{
		engi.NewRegion(texture, 0, 0, width, height),
		engi.NewRegion(texture, 0, height, width, height),
		engi.NewRegion(texture, 0, height*2, width, height),
		engi.NewRegion(texture, 0, height*3, width, height),
	}

	bricks = make([]*Entity, 0)
	for i := 0; i < 7; i++ {
		for j := 0; j < 4; j++ {
			brick := NewEntity(brickRegions[j])
			brick.Position.X = float32(100 + i*100)
			brick.Position.Y = float32(50 + j*75)
			bricks = append(bricks, brick)
		}
	}
}

func (g *Game) BallHitBrick(ball, brick *Entity) {
	score += 10
	scoreText = fmt.Sprintf("score: %d", score)

	// Top
	ballBottom := ball.LastPos.Y + ball.Height()/2
	brickTop := brick.Position.Y - brick.Height()/2
	if ballBottom <= brickTop {
		ball.Velocity.Y = -ball.Velocity.Y
		ball.Position.Y = brickTop - ball.Height()/2
		return
	}

	// Bottom
	ballTop := ball.LastPos.Y - ball.Height()/2
	brickBottom := brick.Position.Y + brick.Height()/2
	if ballTop >= brickBottom {
		ball.Velocity.Y = -ball.Velocity.Y
		ball.Position.Y = brickBottom + ball.Height()/2
		return
	}

	// Left
	ballRight := ball.LastPos.X + ball.Width()/2
	brickLeft := brick.Position.X - brick.Width()/2
	if ballRight <= brickLeft {
		ball.Velocity.X = -ball.Velocity.X
		ball.Position.X = brickLeft - ball.Width()/2
		return
	}

	// Right
	ballLeft := ball.LastPos.X - ball.Width()/2
	brickRight := brick.Position.X + brick.Width()/2
	if ballLeft >= brickRight {
		ball.Velocity.X = -ball.Velocity.X
		ball.Position.X = brickRight + ball.Width()/2
		return
	}
}

func main() {
	engi.Open("Breakout", 800, 600, false, &Game{})
}
