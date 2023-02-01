package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Circle struct {
	xc, yc int
	r      int
	c      color.Color
}

type game struct {
	c Circle
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {

	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	DrawCircle(screen, g.c.xc, g.c.yc, g.c.r, g.c.c)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &game{Circle{screenWidth / 2, screenHeight / 2, 100, color.RGBA{100, 100, 255, 255}}}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func DrawCircle(img *ebiten.Image, xc, yc, r int, c color.Color) {
	x1, y1 := 0, r
	d := 2*(x1+1)*(x1+1) + y1*y1 + (y1-1)*(y1-1) - 2*r*r
	for x, y := x1, y1; x <= y; x++ {
		img.Set(xc+x, yc+y, c) // 2
		img.Set(xc-x, yc+y, c) // 3
		img.Set(xc+x, yc-y, c) // 7
		img.Set(xc-x, yc-y, c) // 6
		img.Set(xc+y, yc+x, c) // 1
		img.Set(xc+y, yc-x, c) // 8
		img.Set(xc-y, yc+x, c) // 4
		img.Set(xc-y, yc-x, c) // 5

		if d <= 0 {
			d += 4*x + 6
		} else {
			d += 4*x - 4*y + 10
			y--
		}
	}
}
