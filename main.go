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
	level *Level
}

// Figure represents a moving and drawable character object in the game
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
	screen.Fill(color.RGBA{0, 255, 255, 255})
	game.fig.Draw(screen)
	game.level.Draw(screen)
}

// Layout determines the canvas size / number of individually drawable pixels
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / 4, outsideHeight / 4
}

func main() {
	game := Game{
		fig:   NewCharacter(Position{80, 60}),
		level: emptyLevel("empty", 160, 128, 16, 16),
	}
	ebiten.SetWindowSize(640, 512)
	ebiten.SetWindowTitle("The Game")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
