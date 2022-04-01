package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
)

const RoundTo float64 = 1000000 // round to 6 decimals

func simulation(n int16, printGrid bool) {
	fmt.Println("Running simulation...")
	var totalCount, i int16 = 0, 0
	for i = 0; i < n; i++ {
		grid := NewGrid()
		grid.RingBell(50)
		totalCount += grid.GetEmptyCellsCount()
		if printGrid {
			grid.Print("")
		}
	}
	avg := float64(totalCount) / float64(n)
	avgCount := math.Round(avg*RoundTo) / RoundTo
	fmt.Printf("Empty cells count for current simulation (%d times) - %f\n", n, avgCount)
}

func parallelSimulation(n int16) {
	var wg sync.WaitGroup
	var totalCount int = 0
	var i int16
	// channel to store individual empty cell counts of each simulation
	counts := make(chan int16, n)

	for i = 0; i < n; i++ {
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
		totalCount += int(c)
	}
	avg := float64(totalCount) / float64(n)
	avgCount := math.Round(avg*RoundTo) / RoundTo
	fmt.Printf("Average empty cells count for parallel simulation (%d times) - %f\n", n, avgCount)
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

	times := flag.Int("times", 1, "Number of times to run simulation.")
	isParallel := flag.Bool("parallel", false, "Determines if simulation should run in parallel.")
	isProfiler := flag.Bool("profile", false, "Determines if simulation should run with profiler enabled.")
	isPrintGrid := flag.Bool("display-grid", false,
		"Determines if grid matrix should be displayed in termial (Supported with simulation in sequence only).")

	flag.Parse()

	logger := initLogs()
	defer logger.Close()

	if *isProfiler {
		prof, perr := os.Create("cpu.pprof")
		if perr != nil {
			log.Fatal(perr)
		}
		pprof.StartCPUProfile(prof)
		defer pprof.StopCPUProfile()
	}

	PrintDebugInfo(func() {
		if *isParallel {
			parallelSimulation(int16(*times))
		} else {
			simulation(int16(*times), *isPrintGrid)
		}
	})
}
