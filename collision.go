package main

type Rectangle struct {
	upperLeft  Position
	lowerRight Position
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
	if r1.lowerRight.x > r2.upperLeft.x &&
		r1.lowerRight.y > r2.upperLeft.y &&
		r2.lowerRight.x > r1.upperLeft.x &&
		r2.lowerRight.y > r1.upperLeft.y {
		return true
	}

	if r2.lowerRight.x > r1.upperLeft.x &&
		r2.lowerRight.y > r1.upperLeft.y &&
		r1.lowerRight.x > r2.upperLeft.x &&
		r1.lowerRight.y > r2.upperLeft.y {
		return true
	}

	return false
}
