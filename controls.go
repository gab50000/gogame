package gogame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func WhichKeyPressed(game *Game) (directions []direction) {
	for _, dirKey := range []ebiten.Key{
		ebiten.KeyUp,
		ebiten.KeyDown,
		ebiten.KeyLeft,
		ebiten.KeyRight,
	} {
		if ebiten.IsKeyPressed(dirKey) {
			switch dirKey {
			case ebiten.KeyUp:
				directions = append(directions, up)
			case ebiten.KeyDown:
				directions = append(directions, down)
			case ebiten.KeyLeft:
				directions = append(directions, left)
			case ebiten.KeyRight:
				directions = append(directions, right)
			}
		}
	}
	return directions
}
