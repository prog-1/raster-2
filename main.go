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

func DrawCircle(img *ebiten.Image, xoffset, yoffset, r int, c color.Color) {
	x1, y1 := 0, r
	d := 2*(x1+1)*(x1+1) + y1*y1 + (y1-1)*(y1-1) - 2*r*r
	for x, y := x1, y1; x <= y; x++ {
		img.Set(xoffset+x, yoffset+y, c) // 2
		img.Set(xoffset-x, yoffset+y, c) // 3
		img.Set(xoffset+x, yoffset-y, c) // 7
		img.Set(xoffset-x, yoffset-y, c) // 6
		img.Set(xoffset+y, yoffset+x, c) // 1
		img.Set(xoffset+y, yoffset-x, c) // 8
		img.Set(xoffset-y, yoffset+x, c) // 4
		img.Set(xoffset-y, yoffset-x, c) // 5

		if d <= 0 {
			d += 4*x + 6
		} else {
			d += 4*x - 4*y + 10
			y--
		}
	}
}

type game struct{}

func (*game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (*game) Update() error                             { return nil }
func (*game) Draw(screen *ebiten.Image) {

	DrawCircle(screen, 300, 300, 100, color.White)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
