package main

import "fmt"

func singeSimulation() {
	grid := NewGrid()
	grid.Print("Grid before simulation")
	grid.RingBell(50)
	grid.Print("Grid after simulation")

	emptyCount := grid.GetEmptyCellsCount()
	fmt.Printf("Empty cells count is - %d", emptyCount)
}

func main() {
	singeSimulation()
}
