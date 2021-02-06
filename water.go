package gogame

type Water struct {
	tiles []*Tile
}

type Move struct {
	posFrom Position
	posTo   Position
	dir     direction
}

func (water Water) getPositions() []Position {
	positions := make([]Position, 0)

	for _, tile := range water.tiles {
		positions = append(positions, tile.bounds.LowerRight)
	}
	return positions
}

func (water Water) getFreeFields() []Position {
	positions := make([]Position, 0)
	return positions

}

func (water Water) getMoves() []Move {
	moves := make([]Move, 1)

	return moves
}

func (water Water) calcEnergy() int {
	energy := 0
	for _, tile := range water.tiles {
		energy += tile.bounds.LowerRight.y
	}
	return energy
}
