package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	_ "net/http/pprof"
	"os"
	"sync"
)

const RoundTo float64 = 1000000 // round to 6 decimals

func simulation(n int) {
	fmt.Println("Running simulation...")
	totalCount := 0
	for i := 0; i < n; i++ {
		grid := NewGrid()
		grid.RingBell(50)
		totalCount += grid.GetEmptyCellsCount()
	}
	avg := float64(totalCount) / float64(n)
	avgCount := math.Round(avg*RoundTo) / RoundTo
	fmt.Printf("Empty cells count for current simulation (%d times) - %f\n", n, avgCount)
}

func parallelSimulation(n int) {
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
	fmt.Printf("Empty cells count for parallel simulation (%d times) - %f\n", n, avgCount)
}

func initLogs() *os.File {
	LOG_FILE := "./debug.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return logFile
}

func main() {

	times := flag.Int("times", 1, "Number of times to run simulation")
	is_parallel := flag.Bool("parallel", false, "Determines if simulation should run in parallel")

	flag.Parse()

	logger := initLogs()
	defer logger.Close()

	var f func(int)

	if *is_parallel {
		f = parallelSimulation
	} else {
		f = simulation
	}

	PrintDebugInfo(func() {
		f(*times)
	})
}
