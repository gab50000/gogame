package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func whichKeyPressed(game *Game) direction {
	for _, dirKey := range []ebiten.Key{
		ebiten.KeyUp,
		ebiten.KeyDown,
		ebiten.KeyLeft,
		ebiten.KeyRight,
	} {
		if ebiten.IsKeyPressed(dirKey) {
			switch dirKey {
			case ebiten.KeyUp:
				return up
			case ebiten.KeyDown:
				return down
			case ebiten.KeyLeft:
				return left
			case ebiten.KeyRight:
				return right
			}
		}
	}
	log.Println("No key pressed")
	return noDirection
}
