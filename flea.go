package main

import (
	"math/rand"
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

// Performs flea move to the neighbor cell
func (flea *Flea) move(allowedDirection []Direction) {
	var nCell *Cell

	// randomly select direction to jump based on allowedDirection
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(len(allowedDirection))
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
