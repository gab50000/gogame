package gogame

// Rectangle is used for collision detection
type Rectangle struct {
	UpperLeft  Position
	LowerRight Position
}

/*
 -------------------------------------> x
| *----------------*
| |                |
| |      r1        |
| |                |
| |                |
| |        *-------|--------*
| |        |       |        |
| *--------|-------*        |
|          |      r2        |
|          |                |
|          *----------------*
V y
*/
func (r1 Rectangle) collidesWith(r2 Rectangle) bool {
	if r1.LowerRight.x > r2.UpperLeft.x &&
		r1.LowerRight.y > r2.UpperLeft.y &&
		r2.LowerRight.x > r1.UpperLeft.x &&
		r2.LowerRight.y > r1.UpperLeft.y {
		return true
	}

	if r2.LowerRight.x > r1.UpperLeft.x &&
		r2.LowerRight.y > r1.UpperLeft.y &&
		r1.LowerRight.x > r2.UpperLeft.x &&
		r1.LowerRight.y > r2.UpperLeft.y {
		return true
	}

	return false
}

func (r1 Rectangle) equals(r2 Rectangle) bool {
	return r1.UpperLeft.equals(r2.UpperLeft) &&
		r1.LowerRight.equals(r2.LowerRight)
}
