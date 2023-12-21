package main

type tile interface {
	determineNextPoints(direction Direction, currentPoint point) map[Direction]point
	energise(direction Direction)
	isEnergised() bool
	hasBeamedInDirection(direction Direction) bool
	reset()
}

type verticalSplinter struct {
	energised        bool
	directionsBeemed []Direction
}

func (t *verticalSplinter) energise(direction Direction) {
	t.energised = true
	t.directionsBeemed = append(t.directionsBeemed, direction)
}

func (t *verticalSplinter) hasBeamedInDirection(direction Direction) bool {
	return findHasBeamedInDirection(t.directionsBeemed, direction)
}

func (t *verticalSplinter) isEnergised() bool {
	return t.energised
}

func (t *verticalSplinter) reset() {
	t.directionsBeemed = []Direction{}
	t.energised = false
}

func (t verticalSplinter) determineNextPoints(direction Direction, currentPoint point) map[Direction]point {
	var results map[Direction]point = map[Direction]point{}
	if direction == East || direction == West {
		results[North] = North.nextPoint(currentPoint)
		results[South] = South.nextPoint(currentPoint)
	} else {
		results[direction] = direction.nextPoint(currentPoint)
	}
	return results
}

type horizontalSplinter struct {
	energised        bool
	directionsBeemed []Direction
}

func (t *horizontalSplinter) energise(direction Direction) {
	t.energised = true
	t.directionsBeemed = append(t.directionsBeemed, direction)
}

func (t *horizontalSplinter) hasBeamedInDirection(direction Direction) bool {
	return findHasBeamedInDirection(t.directionsBeemed, direction)
}

func (t *horizontalSplinter) isEnergised() bool {
	return t.energised
}

func (t *horizontalSplinter) reset() {
	t.directionsBeemed = []Direction{}
	t.energised = false
}

func (t horizontalSplinter) determineNextPoints(direction Direction, currentPoint point) map[Direction]point {
	var results map[Direction]point = map[Direction]point{}
	if direction == North || direction == South {
		results[West] = West.nextPoint(currentPoint)
		results[East] = East.nextPoint(currentPoint)

	} else {
		results[direction] = direction.nextPoint(currentPoint)
	}
	return results
}

type empty struct {
	energised        bool
	directionsBeemed []Direction
}

func (t *empty) energise(direction Direction) {
	t.energised = true
	t.directionsBeemed = append(t.directionsBeemed, direction)
}

func (t *empty) hasBeamedInDirection(direction Direction) bool {
	return findHasBeamedInDirection(t.directionsBeemed, direction)
}

func (t *empty) isEnergised() bool {
	return t.energised
}

func (t *empty) reset() {
	t.directionsBeemed = []Direction{}
	t.energised = false
}

func (t empty) determineNextPoints(direction Direction, currentPoint point) map[Direction]point {
	var results map[Direction]point = map[Direction]point{}
	results[direction] = direction.nextPoint(currentPoint)
	return results
}

type forwardMirror struct {
	energised        bool
	directionsBeemed []Direction
}

func (t *forwardMirror) energise(direction Direction) {
	t.energised = true
	t.directionsBeemed = append(t.directionsBeemed, direction)
}

func (t *forwardMirror) hasBeamedInDirection(direction Direction) bool {
	return findHasBeamedInDirection(t.directionsBeemed, direction)
}

func (t *forwardMirror) isEnergised() bool {
	return t.energised
}

func (t *forwardMirror) reset() {
	t.directionsBeemed = []Direction{}
	t.energised = false
}

func (t forwardMirror) determineNextPoints(direction Direction, currentPoint point) map[Direction]point {
	var results map[Direction]point = map[Direction]point{}
	if direction == East {
		results[North] = North.nextPoint(currentPoint)
	} else if direction == West {
		results[South] = South.nextPoint(currentPoint)
	} else if direction == South {
		results[West] = West.nextPoint(currentPoint)
	} else {
		results[East] = East.nextPoint(currentPoint)
	}
	return results
}

type backwardMirror struct {
	energised        bool
	directionsBeemed []Direction
}

func (t *backwardMirror) energise(direction Direction) {
	t.energised = true
	t.directionsBeemed = append(t.directionsBeemed, direction)
}

func (t *backwardMirror) hasBeamedInDirection(direction Direction) bool {
	return findHasBeamedInDirection(t.directionsBeemed, direction)
}

func (t *backwardMirror) isEnergised() bool {
	return t.energised
}

func (t *backwardMirror) reset() {
	t.directionsBeemed = []Direction{}
	t.energised = false
}

func (t backwardMirror) determineNextPoints(direction Direction, currentPoint point) map[Direction]point {
	var results map[Direction]point = map[Direction]point{}
	if direction == East {
		results[South] = South.nextPoint(currentPoint)
	} else if direction == West {
		results[North] = North.nextPoint(currentPoint)
	} else if direction == South {
		results[East] = East.nextPoint(currentPoint)
	} else {
		results[West] = West.nextPoint(currentPoint)
	}
	return results
}

func findHasBeamedInDirection(directionsBeamed []Direction, direction Direction) bool {
	for i := 0; i < len(directionsBeamed); i++ {
		if directionsBeamed[i] == direction {
			return true
		}
	}
	return false
}
