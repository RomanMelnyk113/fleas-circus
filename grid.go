package main

import (
	"fmt"
	"strings"
)

const GridSize = 30 // can be dynamic but according to the task should be 30
const FleasCount = GridSize * GridSize

type Grid struct {
	Matrix [GridSize][GridSize]Cell
	Fleas  [FleasCount]*Flea
}

// Initialize default state of grid where each flea is on separate cell
func NewGrid() *Grid {
	k := 0
	g := &Grid{}
	g.walk(func(i int, j int) {
		cell, flea := NewCell(i, j, g)
		g.Matrix[i][j] = *cell
		g.Fleas[k] = flea
		k++
	})

	return g
}

// Iterate over grid matrix and run custom func
func (g *Grid) walk(f func(int, int)) {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			f(i, j)
		}
	}
}

// Print matrix with numbers of fleas inside each cell
func (g *Grid) Print() {
	g.walk(func(i int, j int) {
		fmt.Print(len(g.Matrix[i][j].Fleas))
		if j == GridSize-1 {
			fmt.Println()
		}
	})
	fmt.Println(strings.Repeat("=", GridSize))
}

func (g *Grid) GetEmptyCellsCount() int {
	emptyCount := 0
	g.walk(func(i int, j int) {
		if len(g.Matrix[i][j].Fleas) == 0 {
			emptyCount++
		}
	})
	return emptyCount
}

// Simulates singe bell ring
func (g *Grid) RingBell(n int) {
	for i := 0; i < n; i++ {
		for _, f := range g.Fleas {
			f.Jump()
		}
		fmt.Printf("Empty cells: %d\n", g.GetEmptyCellsCount())
	}
}
