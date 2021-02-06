package gogame

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
	options.GeoM.Translate(float64(tile.bounds.UpperLeft.x), float64(tile.bounds.UpperLeft.y))
	screen.DrawImage(tile.img, &options)
}

func NewColoredTile(pos Position, width, height int, color color.Color) *Tile {
	blockImg := NewColoredImage(width, height, color)
	return &Tile{
		bounds: Rectangle{
			Position{pos.x, pos.y},
			Position{pos.x + width, pos.y + height},
		},
		img: blockImg,
	}
}

func NewRandomImage(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})
		}
	}

	return img
}

func NewColoredImage(width, height int, color color.Color) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	img.Fill(color)
	return img
}

func NewColoredImageWithFeatures(width, height, nFeatures int, background, foreground color.Color) *ebiten.Image {
	img := NewColoredImage(width, height, background)

	for n := 0; n < nFeatures; n++ {
		i := rand.Intn(height)
		j := rand.Intn(width)

		img.Set(j, i, foreground)
		img.Set(j, (i+1)%height, foreground)
	}
	return img
}

func NewTileFromImage(pos Position, width, height int, image *ebiten.Image) *Tile {
	return &Tile{
		bounds: Rectangle{
			pos,
			Position{pos.x + width, pos.y + height},
		},
		img: image,
	}
}

func NewRandomTile(pos Position, width, height int, color color.Color) *Tile {
	randImg := NewRandomImage(width, height)
	return NewTileFromImage(pos, width, height, randImg)
}

func NewGrassTile(pos Position, width, height int) *Tile {
	img := NewColoredImageWithFeatures(
		width, height, 50, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 200, 0, 255})
	return NewTileFromImage(pos, width, height, img)
}

func NewEarthTile(pos Position, width, height int) *Tile {
	img := NewColoredImageWithFeatures(
		width, height, 50, color.RGBA{95, 46, 0, 255}, color.RGBA{89, 35, 0, 255})
	return NewTileFromImage(pos, width, height, img)
}

func NewWaterTile(pos Position, width, height int) *Tile {
	img := NewColoredImage(width, height, color.RGBA{0, 0, 255, 255})
	return NewTileFromImage(pos, width, height, img)
}
