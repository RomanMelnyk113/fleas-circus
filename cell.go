package main

type Cell struct {
	Fleas  map[int]*Flea
	Left   *Cell
	Right  *Cell
	Top    *Cell
	Bottom *Cell
	X      int
	Y      int
}

// Initialize new cell with 1 flea inside
func NewCell(x int, y int, grid *Grid) (*Cell, *Flea) {
	cell := &Cell{X: x, Y: y, Fleas: make(map[int]*Flea)}
	fleaID := x*GridSize + y // get unique ID
	flea := &Flea{ID: fleaID, Cell: cell}
	cell.Fleas[flea.ID] = flea

	// fill neighbor cells
	if y > 0 {
		cell.Top = &grid.Matrix[x][y-1]
	}
	if y < GridSize-1 {
		cell.Bottom = &grid.Matrix[x][y+1]
	}
	if x > 0 {
		cell.Left = &grid.Matrix[x-1][y]
	}
	if x < GridSize-1 {
		cell.Right = &grid.Matrix[x+1][y]
	}
	return cell, flea
}
