package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	img   *ebiten.Image
	count int
)

type Game struct{}

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
	count++
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 50)
	op.GeoM.Rotate(float64(count%360) * 3.14159 / 180)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Rotate")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
