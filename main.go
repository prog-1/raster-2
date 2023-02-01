package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Game struct {
		r      int
		xc, yc int
		t      bool
	}
)

var col = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

const (
	winTitle            = "raster"
	winWidth, winHeight = 300, 300
)

func (g *Game) DrawCircle(img *ebiten.Image) {
	x1, y1 := 0, g.r
	d := 2*(x1+1)*(x1+1) + y1*y1 + (y1-1)*(y1-1) - 2*g.r*g.r
	for x, y := x1, y1; x <= y; x++ {
		img.Set(g.xc+x, g.yc+y, col)
		img.Set(g.xc+x, g.yc-y, col)
		img.Set(g.xc+y, g.yc+x, col)
		img.Set(g.xc+y, g.yc-x, col)
		img.Set(g.xc-x, g.yc+y, col)
		img.Set(g.xc-x, g.yc-y, col)
		img.Set(g.xc-y, g.yc+x, col)
		img.Set(g.xc-y, g.yc-x, col)
		if d <= 0 {
			d += 4*x + 6
		} else {
			d += 4*(x-y) + 10
			y--
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawCircle(screen)

}
func (g *Game) rad() {
	if g.xc+g.r < winWidth && g.xc-g.r > 0 && g.yc+g.r < winHeight && g.yc-g.r > 0 && g.t {
		g.r++
		return
	} else if g.t {
		g.t = !g.t
		g.r--
		return
	} else if g.r > 0 {
		g.r--
		return
	}
	g.t = !g.t

}
func (g *Game) Update() error {
	g.rad()
	return nil

}

func (g *Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{r: winHeight * 0.3, xc: winWidth / 2, yc: winHeight / 2, t: true}); err != nil {
		log.Fatal(err)
	}
}
