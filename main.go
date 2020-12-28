package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is game
type Game struct {
	count int
	img   *ebiten.Image
}

func (game *Game) Update() error {
	game.count++
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	screen.DrawImage(game.img.SubImage(image.Rect(0, 0, 10, 10)).(*ebiten.Image), &options)

}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	file, err := ebitenutil.OpenFile("sprites.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	game := Game{img: ebiten.NewImageFromImage(img)}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
