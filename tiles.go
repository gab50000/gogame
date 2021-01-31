package main

import (
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
	blockImg := newColoredImage(width, height, color)
	return &Tile{
		bounds: Rectangle{
			Position{pos.x, pos.y},
			Position{pos.x + width, pos.y + height},
		},
		img: blockImg,
	}
}

func newRandomImage(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})
		}
	}

	return img
}

func newColoredImage(width, height int, color color.Color) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	img.Fill(color)
	return img
}

func newColoredImageWithFeatures(width, height, nFeatures int, background, foreground color.Color) *ebiten.Image {
	img := newColoredImage(width, height, background)

	for n := 0; n < nFeatures; n++ {
		i := rand.Intn(height)
		j := rand.Intn(width)

		img.Set(j, i, foreground)
		img.Set(j, (i+1)%height, foreground)
	}
	return img
}

func newTileFromImage(pos Position, width, height int, image *ebiten.Image) *Tile {
	return &Tile{
		bounds: Rectangle{
			pos,
			Position{pos.x + width, pos.y + height},
		},
		img: image,
	}
}

func newRandomTile(pos Position, width, height int, color color.Color) *Tile {
	randImg := newRandomImage(width, height)
	return newTileFromImage(pos, width, height, randImg)
}

func newGrassTile(pos Position, width, height int) *Tile {
	img := newColoredImageWithFeatures(
		width, height, 50, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 200, 0, 255})
	return newTileFromImage(pos, width, height, img)
}

func newEarthTile(pos Position, width, height int) *Tile {
	img := newColoredImageWithFeatures(
		width, height, 50, color.RGBA{95, 46, 0, 255}, color.RGBA{89, 35, 0, 255})
	return newTileFromImage(pos, width, height, img)
}
