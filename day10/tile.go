package main

type tile interface {
	getConnectionPoints() []Direction
	// connects(tile) bool
}

type verticalPipe struct {
}

func (t verticalPipe) getConnectionPoints() []Direction {
	return []Direction{North, South}
}

type horizontalPipe struct {
}

func (t horizontalPipe) getConnectionPoints() []Direction {
	return []Direction{East, West}
}

type lBend struct {
}

func (t lBend) getConnectionPoints() []Direction {
	return []Direction{North, East}
}

type jBend struct {
}

func (t jBend) getConnectionPoints() []Direction {
	return []Direction{North, West}
}

type sevenBend struct {
}

func (t sevenBend) getConnectionPoints() []Direction {
	return []Direction{South, West}
}

type fBend struct {
}

func (t fBend) getConnectionPoints() []Direction {
	return []Direction{South, East}
}

type ground struct {
}

func (t ground) getConnectionPoints() []Direction {
	return []Direction{}
}

type startPipe struct {
}

func (t startPipe) getConnectionPoints() []Direction {
	return []Direction{North, East, South, West}
}

func connectsInDirection(to tile, direction Direction) bool {
	requiredConnectionPoint := direction.opposite()
	var connectionPoints []Direction = to.getConnectionPoints()
	var connects bool = false
	for _, connectionPoint := range connectionPoints {
		if connectionPoint == requiredConnectionPoint {
			connects = true
			break
		}
	}
	return connects
}
