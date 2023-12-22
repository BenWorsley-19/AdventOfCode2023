package main

type Point struct {
	X int64
	Y int64
}

type Direction int

const (
	North Direction = iota + 1
	East
	South
	West
	Stationary
)

var directions []Direction = []Direction{North, East, South, West}

func (d Direction) Opposite() Direction {
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

func (d Direction) NextPoint(currentPoint Point, stepsInDir int64) Point {
	var x int64
	var y int64
	if d == North {
		x = stepsInDir * -1
		y = stepsInDir * 0
	} else if d == East {
		x = stepsInDir * 0
		y = stepsInDir * 1
	} else if d == South {
		x = stepsInDir * 1
		y = stepsInDir * 0
	} else {
		x = stepsInDir * 0
		y = stepsInDir * -1
	}
	return Point{
		X: currentPoint.X + x,
		Y: currentPoint.Y + y,
	}
}
