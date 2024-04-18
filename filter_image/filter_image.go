package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("ebiten.png")
	if err != nil {
		log.Fatal(err)
	}
}

func NewGame() *Game {
	g := &Game{}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	screen.DrawImage(img, op)

	// draw the image with a filter.
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 50)
	op.GeoM.Scale(2, 2)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(img, op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Filter Image")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
