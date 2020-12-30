package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game is game
type Game struct {
	count int
	theta float64
	fig   *Character
}

// Figure bla
type Figure interface {
	Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions)
	Move(dx, dy int)
}

// Update bla
func (game *Game) Update() error {
	dir := whichKeyPressed(game)
	game.count++
	game.fig.Update(dir)
	return nil
}

// Draw bla
func (game *Game) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	screen.Fill(color.RGBA{0, 255, 255, 255})
	options.GeoM.Translate(float64(screen.Bounds().Max.X/2), float64(screen.Bounds().Max.Y/2))
	game.fig.Draw(screen, &options)
}

// Layout bla
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / 4, outsideHeight / 4
}

func main() {
	game := Game{fig: NewCharacter()}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("The Game")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
