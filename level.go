package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tile represents a level element
type Tile struct {
	position Position
	img      *ebiten.Image
}

type Level struct {
	name  string
	tiles []Tile
}

func (tile *Tile) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	blockImg := ebiten.NewImage(10, 10)
	options.GeoM.Translate(float64(tile.position.x), float64(tile.position.y))
	blockImg.Fill(color.Black)
	screen.DrawImage(blockImg, &options)
}

func (level *Level) Draw(screen *ebiten.Image) {
	for _, tile := range level.tiles {
		tile.Draw(screen)
	}
}
