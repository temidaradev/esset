package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/temidaradev/esset/example/assets"
)

type Char struct {
	x  int
	y  int
	vx int
	vy int
}

const (
	groundY = 395
	unit    = 10
)

func (c *Char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *Char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 2
	} else if c.vx < 0 {
		c.vx += 2
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

type Player struct {
	player *Char
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &Char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = 5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = -5 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.player.tryJump()
	}

	p.player.update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image, g Game) {
	s := assets.Idle
	if p.player.vx > 0 {
		s = assets.Right
	} else if p.player.vx < 0 {
		s = assets.Left
	}

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(0.3, 0.3)
	op2.GeoM.Translate(float64(p.player.x)/unit, float64(p.player.y)/unit)
	screen.DrawImage(s, op2)
}

type Game struct {
	player Player
}

const (
	sW = 635
	sH = 475
)

var (
	LightBlue = color.RGBA{0x80, 0xa0, 0xc0, 0xff}
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(LightBlue)
	g.player.Draw(screen, *g)

}

func (g *Game) Update() error {
	g.player.Update()

	vsync := ebiten.IsVsyncEnabled()

	if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		ebiten.SetVsyncEnabled(!vsync)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.player.x = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.player.x = 0
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sW, sH
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
