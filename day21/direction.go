package main

type point struct {
	x int64
	y int64
}

type direction int

const (
	North direction = iota + 1
	East
	South
	West
	Stationary
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
	if d == Stationary {
		return Stationary
	}
	return West
}

func (d direction) nextPoint(currentPoint point, stepsInDir int64) point {
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
	return point{
		x: currentPoint.x + x,
		y: currentPoint.y + y,
	}
}
