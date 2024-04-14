package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	// give the "relative" path to the image.
	img, _, err = ebitenutil.NewImageFromFile("ebiten.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

// NewGame creates a new Game Object and initializes the data.
func NewGame() *Game {
	g := &Game{}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render Image")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
