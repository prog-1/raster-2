package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func DrawCircle(img *ebiten.Image, x, y, r int, c color.Color) {
	d := 3 - 2*r
	for x2, y2 := 0, r; x2 <= y2; x2++ {
		img.Set(x2+x, -y2+y, c)                                         //1            8    1
		img.Set(y2+x, -x2+y, color.RGBA{R: 255, G: 130, B: 0, A: 1})    //2        7            2
		img.Set(y2+x, x2+y, color.RGBA{R: 255, G: 242, B: 0, A: 1})     //3
		img.Set(x2+x, y2+y, color.RGBA{R: 8, G: 255, B: 0, A: 1})       //4        6            3
		img.Set(-x2+x, y2+y, color.RGBA{R: 0, G: 255, B: 252, A: 1})    //5            5     4
		img.Set(-y2+x, x2+y, color.RGBA{R: 0, G: 77, B: 255, A: 1})     //6
		img.Set(-y2+x, -x2+y, color.RGBA{R: 151, G: 0, B: 255, A: 1})   //7
		img.Set(-x2+x, -y2+y, color.RGBA{R: 255, G: 255, B: 255, A: 1}) //8
		if d <= 0 {
			d += 4*x2 + 6
		} else {
			d += 4*x2 - 4*y2 + 10
			y2--
		}
	}
}

type Game struct {
	x, y, r int
}

func (g *Game) Update() error {
	if g.y+g.r < screenHeight-20 {
		g.r++
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawCircle(screen, g.x, g.y, g.r, color.RGBA{R: 255, G: 0, B: 0, A: 1})
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{x: 320, y: 240, r: 100}); err != nil {
		log.Fatal(err)
	}
}
