package main

type point struct {
	x int
	y int
}

type direction int

const (
	North direction = iota + 1
	East
	South
	West
	None
)

var directions []direction = []direction{North, East, South, West}

func (d direction) Opposite() direction {
	if d == North {
		return South
	}
	if d == South {
		return North
	}
	if d == West {
		return East
	}
	if d == None {
		return None
	}
	return West
}

func (d direction) nextPoint(currentPoint point) point {
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
