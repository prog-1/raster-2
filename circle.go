package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//---------------------------Declaration--------------------------------

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	//here all the global variables are stored
	width, height int     //screen size
	c             *circle //circle struct
}

type circle struct {
	x, y, r int //center coordinates and radius
	color   color.RGBA
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update

	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {

	//Circle Draw
	g.DrawCircle(screen, g.c.x, g.c.y, g.c.r, g.c.color)
}

//-------------------------Functions----------------------------------

func (g *Game) DrawCircle(screen *ebiten.Image, cx, cy, r int, c color.Color) {
	//screen.Set(cx, cy, c) //center point

	d := 3 - 2*r //simplified initial decision parameter (d0) formula

	for x, y := 0, r; x <= y; x++ {
		//filling octants (with y reversed)
		screen.Set(cx+x, cy+y, c) //1st octant
		screen.Set(cx+y, cy+x, c) //2nt octant
		screen.Set(cx-y, cy+x, c) //3rt octant
		screen.Set(cx-x, cy+y, c) //4th octant
		screen.Set(cx-x, cy-y, c) //5th octant
		screen.Set(cx-y, cy-x, c) //6th octant
		screen.Set(cx+y, cy-x, c) //7th octant
		screen.Set(cx+x, cy-y, c) //8th octant

		if d < 0 { //if circle is closer to outside point
			d = d + 4*x + 6 //simplified f(x+1,y)
		} else { //if circle is closer to inside point
			d = d + 4*(x-y) + 10 //simplified f(x+1,y-1)
			y--
		}
	}

}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Circle")

	//creating game instance
	g := &Game{width: screenWidth, height: screenHeight,
		c: &circle{x: screenWidth / 2, y: screenHeight / 2, r: 100, color: color.RGBA{255, 255, 255, 255}}} //declaring circle

	//running game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
