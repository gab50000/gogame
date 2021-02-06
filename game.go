package gogame

import (
	"image/color"

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
	dirs := WhichKeyPressed(game)
	game.count++
	game.fig = game.fig.Update(dirs, game.level)
	return nil
}

// Draw draws the game including the level and all characters
func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 255, 255})
	game.fig.drawBounds(screen)
	game.level.Draw(screen)
}

// Layout determines the canvas size / number of individually drawable pixels
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / 2, outsideHeight / 2
}
