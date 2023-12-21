package main

type Direction int

const (
	North Direction = iota + 1
	East
	South
	West
	Stationary
)

func (d Direction) opposite() Direction {
	if d == North {
		return South
	}
	if d == South {
		return North
	}
	if d == West {
		return East
	}
	if d == Stationary {
		return Stationary
	}
	return West
}

func (d Direction) nextPoint(currentPoint point) point {
	var x int
	var y int
	if d == North {
		x = -1
		y = 0
	} else if d == East {
		x = 0
		y = 1
	} else if d == South {
		x = 1
		y = 0
	} else {
		x = 0
		y = -1
	}
	return point{
		x: currentPoint.x + x,
		y: currentPoint.y + y,
	}
}

// var directions []point = []point{
// 	point{x: -1, y: 0},
// 	point{x: 1, y: 0},
// 	point{x: 0, y: -1},
// 	point{x: 0, y: 1},
// }

// func up() point {
// 	return point{x: -1, y: 0}
// }

// func right() point {
// 	return point{x: 1, y: 0}
// }

// func down() point {
// 	return point{x: 0, y: -1}
// }

// func left() point {
// 	return point{x: 0, y: 1}
// }
