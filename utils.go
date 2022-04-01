package main

import (
	"fmt"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Print("Memory usage details: ")
	fmt.Printf("Alloc = %v MiB (%v B)", bToMb(m.Alloc), m.Alloc)
	fmt.Printf("\tTotalAlloc = %v MiB (%v B)", bToMb(m.TotalAlloc), m.TotalAlloc)
	fmt.Printf("\tSys = %v MiB (%v B)", bToMb(m.Sys), m.Sys)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Println()
}

// Measure function execution time
func getExecTime(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
	PrintMemUsage()
}
