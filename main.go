package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y int
}

type Game struct {
	width, height, r, dir, x, y int
	vel Point
}

func DrawCircle(img *ebiten.Image, x, y, r int, c color.Color) {
	d := 3-2*r
	for xi, yi := 0, r; xi <= yi; xi = xi + 1 {
		img.Set(xi+x, yi+y, c)
		img.Set(-xi+x, yi+y, c)
		img.Set(xi+x, -yi+y, c)
		img.Set(-xi+x, -yi+y, c)
		img.Set(yi+x, xi+y, c)
		img.Set(-yi+x, xi+y, c)
		img.Set(yi+x, -xi+y, c)
		img.Set(-yi+x, -xi+y, c)
		if d > 0 {
			d += 4*xi-4*yi+10
			yi--
		} else {
			d += 4*xi+6
		}
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		dir: 1,
		x: width/2,
		y: height/2,
		vel: Point{x: 1, y: 1},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.r += g.dir
	if g.r > 100 {
		g.dir = -1
	} else if g.r < 1 {
		g.dir = 1
	}
	g.x += g.vel.x
	g.y += g.vel.y
	switch {
	case g.x+g.r >= g.width:
		g.x = g.width - g.r - 1
		g.vel.x = -g.vel.x
	case g.x-g.r <= 1:
		g.x = g.r
		g.vel.x = -g.vel.x
	case g.y+g.r >= g.height:
		g.y = g.height - g.r - 1
		g.vel.y = -g.vel.y
	case g.y-g.r <= 1:
		g.y = g.r
		g.vel.y = -g.vel.y
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawCircle(screen, g.x, g.y, g.r, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}