package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game is game
type Game struct {
	count   int
	img     *ebiten.Image
	theta   float64
	figures []Figure
}

type Figure interface {
	Draw(screen *ebiten.Image)
	Move(dx, dy int)
}

func (game *Game) Update() error {
	game.count++
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	screen.Fill(color.RGBA{0, 255, 255, 255})
	x, y := 1, 11
	dx, dy := 15, 15
	subimg := game.img.SubImage(image.Rect(x, y, x+dx, y+dy)).(*ebiten.Image)
	minX, minY := subimg.Bounds().Min.X, subimg.Bounds().Min.Y
	maxX, maxY := subimg.Bounds().Max.X, subimg.Bounds().Max.Y
	centerX, centerY := (minX+maxX)/2, (minY+maxY)/2
	options.GeoM.Translate(-float64(centerX), -float64(centerY))
	options.GeoM.Scale(float64(math.Pow(-1, float64(game.count))), 1)
	options.GeoM.Translate(float64(centerX), float64(centerY))
	screen.DrawImage(subimg, &options)
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
	ebiten.SetWindowTitle("The Game")
	ebiten.SetMaxTPS(10)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
