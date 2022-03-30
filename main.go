package main

import (
	"fmt"
	"math/rand"
	"time"
)

const XCells int = 5
const YCells int = 5

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

type Cell struct {
	Fleas  map[int]*Flea
	Left   *Cell
	Right  *Cell
	Top    *Cell
	Bottom *Cell
	X      int
	Y      int
}

//type Grid struct {
//CellID int
//X      int
//Y      int
//}

const fleasCount = XCells * YCells

var grid [XCells][YCells]Cell
var fleas [fleasCount]*Flea

func newCell(fleas map[int]*Flea, x int, y int) *Cell {
	cell := Cell{Fleas: fleas, X: x, Y: y}

	// fill related cells
	if y > 0 {
		cell.Top = &grid[x][y-1]
	}
	if y < YCells-1 {
		cell.Bottom = &grid[x][y+1]
	}
	if x > 0 {
		cell.Left = &grid[x][x-1]
	}
	if x < XCells-1 {
		cell.Right = &grid[x][x+1]
	}
	return &cell
}

func walkGrid(f func(int, int)) {
	for i := 0; i < XCells; i++ {
		for j := 0; j < YCells; j++ {
			f(i, j)
		}
	}
}

// Initialize default state of grid where each flea is on separate cell
func initGrid() {
	k := 0
	walkGrid(func(i int, j int) {
		f := &Flea{ID: k}
		f.Cell = newCell(map[int]*Flea{k: f}, i, j)
		grid[i][j] = *f.Cell
		fleas[k] = f
		k++
	})
}

func printGrid() {
	fmt.Println("=============")
	walkGrid(func(i int, j int) {
		fmt.Print(len(grid[i][j].Fleas))
		if j == YCells-1 {
			fmt.Println()
		}
	})
}

func move(flea *Flea, allowedDirection []Direction) {
	var nCell *Cell
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
		panic("AAAAAAAAAAAAAAAAAAAAA")
	}
	delete(flea.Cell.Fleas, flea.ID)
	nCell.Fleas[flea.ID] = flea
	flea.Cell = nCell
}

// Perform single flea jump
func jump(flea *Flea) {
	if flea.Cell.X == 0 && flea.Cell.Y == 0 {
		move(flea, []Direction{RIGHT, DOWN})
	} else if flea.Cell.X == XCells-1 && flea.Cell.Y == YCells-1 {
		move(flea, []Direction{LEFT, UP})
	} else if flea.Cell.X == 0 && flea.Cell.Y == YCells-1 {
		move(flea, []Direction{RIGHT, UP})
	} else if flea.Cell.X == XCells-1 && flea.Cell.Y == 0 {
		move(flea, []Direction{LEFT, DOWN})
	} else if flea.Cell.X == 0 {
		move(flea, []Direction{RIGHT, DOWN, UP})
	} else if flea.Cell.Y == 0 {
		move(flea, []Direction{RIGHT, DOWN, LEFT})
	} else if flea.Cell.X == XCells-1 {
		move(flea, []Direction{DOWN, UP, LEFT})
	} else if flea.Cell.Y == YCells-1 {
		move(flea, []Direction{UP, LEFT, RIGHT})
	} else {
		move(flea, []Direction{RIGHT, DOWN, UP, LEFT})
	}
}

// Simulates singe bell ring
func ringBell(n int) {
	for i := 0; i < n; i++ {
		for _, f := range fleas {
			jump(f)
		}
	}
}

func main() {
	initGrid()
	printGrid()
	ringBell(1)
	printGrid()
}
