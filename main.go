package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1290
	screenHeight = 960
	radius       = 69
)

type Game struct {
	width, height int
	last          time.Time
}

func (g *Game) reflections(x, y int, img *ebiten.Image, c color.RGBA) {
	w, h := g.width/2, g.height/2
	img.Set(x+w, y+h, c)
	img.Set(-x+w, y+h, c)
	img.Set(x+w, -y+h, c)
	img.Set(-x+w, -y+h, c)
	img.Set(y+w, x+h, c)
	img.Set(-y+w, x+h, c)
	img.Set(y+w, -x+h, c)
	img.Set(-y+w, -x+h, c)
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) DrawCircle(screen *ebiten.Image) {
	d := 3 - 2*radius
	for x, y := 0, radius; x <= y; x = x + 1 {
		g.reflections(x, y, screen, color.RGBA{255, 0, 0, 255})
		if d > 0 {
			d += 4*x - 4*y + 10
			y--
		} else {
			d += 4*x + 6
		}
	}
}
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawCircle(screen)
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		last:   time.Now(),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
