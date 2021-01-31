package main

import (
	"bufio"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Level contains the level information
type Level struct {
	name  string
	tiles []*Tile
}

// Draw draws the level to the screen
func (level *Level) Draw(screen *ebiten.Image) {
	for _, tile := range level.tiles {
		tile.Draw(screen)
	}
}

func createLevel(filename string) []Tile {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, c := range line {
			if string(c) == "*" {
				log.Print(i, j)
			}
		}
	}

	return []Tile{}
}

func emptyLevel(levelName string, width, height, blockWidth, blockHeight int) *Level {
	tiles := []*Tile{}

	for i := 0; i < width/blockWidth; i++ {
		tiles = append(
			tiles,
			newGrassTile(Position{i * blockWidth, 0}, blockWidth, blockHeight),
			newGrassTile(Position{i * blockWidth, height - blockHeight}, blockWidth, blockHeight),
		)
	}

	for j := 1; j < height/blockHeight-1; j++ {
		tiles = append(
			tiles,
			newColoredTile(Position{0, j * blockHeight}, blockWidth, blockHeight, color.Black),
			newColoredTile(Position{width - blockWidth, j * blockHeight}, blockWidth, blockHeight, color.Black))
	}

	return &Level{levelName, tiles}
}
