package main

import (
	"time"
)

type Direction int

const (
	LEFT  Direction = 0
	RIGHT Direction = 1
	UP    Direction = 2
	DOWN  Direction = 3
)

type Flea struct {
	ID   int
	Cell *Cell
}

// Returns, as an int, a non-negative pseudo-random number in [0,3].
func randNum(n int) int {
	if n > 4 {
		panic("Support only range from 1 to 4")
	}
	// get current timestamp as base for randomizer
	timestamp := time.Now().UnixNano()

	// get last digit from timestamp as rand number
	randNum := int(timestamp % 10)

	// calculate step to check (e.g. n=4 then step = 2 (10/4))
	// so loop below will be running within 4 ranges: [0,2], [2,4], [4,6], [6,8]
	step := 10 / n
	border := 0
	for i := 0; i < n; i++ {
		if randNum >= border && randNum <= step+border {
			return i
		}
		border += step
	}
	// fallback behaviour when randNum is higher than top range border
	// e.g randNum = 9 but top border is 8
	return n - 1
}

// Performs flea move to the neighbor cell
func (flea *Flea) move(allowedDirection []Direction) {
	var nCell *Cell

	// randomly select direction to jump based on allowedDirection
	i := randNum(len(allowedDirection))
	direction := allowedDirection[i]

	switch direction {
	case LEFT:
		nCell = flea.Cell.Left
	case RIGHT:
		nCell = flea.Cell.Right
	case UP:
		nCell = flea.Cell.Top
	case DOWN:
		nCell = flea.Cell.Bottom
	}

	if nCell == nil {
		panic(`Oops! Looks like data provided in "allowedDirection" contains wrong values.
			Make sure correct direction is passed for current flea.`)
	}

	// remove flea from current cell and place it to the next one
	delete(flea.Cell.Fleas, flea.ID)
	flea.Cell = nCell
	flea.Cell.Fleas[flea.ID] = flea
}

// Handles flea logic of jumping to the another cell based on its position on the grid
func (flea *Flea) Jump() {
	if flea.Cell.X == 0 && flea.Cell.Y == 0 {
		flea.move([]Direction{RIGHT, DOWN})
	} else if flea.Cell.X == GridSize-1 && flea.Cell.Y == GridSize-1 {
		flea.move([]Direction{LEFT, UP})
	} else if flea.Cell.X == 0 && flea.Cell.Y == GridSize-1 {
		flea.move([]Direction{RIGHT, UP})
	} else if flea.Cell.X == GridSize-1 && flea.Cell.Y == 0 {
		flea.move([]Direction{LEFT, DOWN})
	} else if flea.Cell.X == 0 {
		flea.move([]Direction{RIGHT, DOWN, UP})
	} else if flea.Cell.Y == 0 {
		flea.move([]Direction{RIGHT, DOWN, LEFT})
	} else if flea.Cell.X == GridSize-1 {
		flea.move([]Direction{DOWN, UP, LEFT})
	} else if flea.Cell.Y == GridSize-1 {
		flea.move([]Direction{UP, LEFT, RIGHT})
	} else {
		flea.move([]Direction{RIGHT, DOWN, UP, LEFT})
	}
}
