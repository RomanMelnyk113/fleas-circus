package main

import (
	"fmt"
	"log"
	"math"
	_ "net/http/pprof"
	"os"
	"sync"
)

const RoundTo float64 = 1000000 // round to 6 decimals

func singeSimulation() {
	fmt.Println("Running single simulation...")
	grid := NewGrid()
	grid.RingBell(50)
	fmt.Printf("Empty cells count is - %d\n", grid.GetEmptyCellsCount())
}

func multipleSimulation(n int) {
	fmt.Println("Running multiple simulation...")
	totalCount := 0
	for i := 0; i < n; i++ {
		grid := NewGrid()
		grid.RingBell(50)
		totalCount += grid.GetEmptyCellsCount()
	}
	avg := float64(totalCount) / float64(n)
	avgCount := math.Round(avg*RoundTo) / RoundTo
	fmt.Printf("Total count: %d\n", totalCount)
	fmt.Printf("Empty cells count for multiple simulation (%d times) - %f\n", n, avgCount)
}

func parallelMultipleSimulation(n int) {
	var wg sync.WaitGroup
	totalCount := 0

	// channel to store individual empty cell counts of each simulation
	counts := make(chan int, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			grid := NewGrid()
			grid.RingBell(50)
			count := grid.GetEmptyCellsCount()
			counts <- count

			wg.Done()
		}()
	}

	wg.Wait() // wait until all paraller simulation finish

	close(counts)

	// iterate within already received empty cell count for each simulation
	// and calculate total empty cells count
	for c := range counts {
		totalCount += c
	}
	avg := float64(totalCount) / float64(n)
	avgCount := math.Round(avg*RoundTo) / RoundTo
	fmt.Printf("Empty cells count for multiple simulation (%d times) - %f\n", n, avgCount)
}

func initLogs() *os.File {
	LOG_FILE := "./debug.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return logFile
}

func main() {
	logger := initLogs()
	defer logger.Close()
	//getExecTime(func() {
	//singeSimulation()
	//})
	//getExecTime(func() {
	//multipleSimulation(30)
	//})
	PrintDebugInfo(func() {
		parallelMultipleSimulation(30)
	})
}
