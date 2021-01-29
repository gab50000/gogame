package main

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tile represents a level element
type Tile struct {
	bounds Rectangle
	img    *ebiten.Image
}

// Draw draws the tile to the screen
func (tile *Tile) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(tile.bounds.upperLeft.x), float64(tile.bounds.upperLeft.y))
	screen.DrawImage(tile.img, &options)
}

func newColoredTile(pos Position, width, height int, color color.Color) *Tile {
	blockImg := ebiten.NewImage(width, height)
	blockImg.Fill(color)
	return &Tile{
		bounds: Rectangle{
			Position{pos.x, pos.y},
			Position{pos.x + width, pos.y + height},
		},
		img: blockImg,
	}
}

func newRandomTile(pos Position, width, height int, color color.Color) *Tile {
	randImg := newRandomImage(width, height)
	blockImg := ebiten.NewImageFromImage(randImg)
	return &Tile{
		bounds: Rectangle{
			Position{pos.x, pos.y},
			Position{pos.x + width, pos.y + height},
		},
		img: blockImg,
	}
}

func newRandomImage(width, height int) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})
		}
	}

	return img
}

func newTileFromImage(pos Position, width, height int, image *image.Image) *Tile {
	return &Tile{}
}
