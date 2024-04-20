package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var (
	img   *ebiten.Image
	count int
)

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
	// 画像サイズを取得
	s := img.Bounds().Size()
	// 画像の中心が原点となるように移動する
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	op.GeoM.Rotate(float64(count%360) * 3.14159 / 180)
	// 回転後に元の位置に戻すことで、画像の中心を軸に回転する
	op.GeoM.Translate(float64(s.X)/2, float64(s.Y)/2+50)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Rotate Image")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
