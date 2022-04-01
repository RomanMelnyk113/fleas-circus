package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Measure function execution time
func PrintDebugInfo(f func()) {
	var m runtime.MemStats

	start := time.Now()
	f()
	elapsed := time.Since(start)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	details := "Resources usage: Time = %v; Memory = %v MiB (%v B); NumGC = %v"
	runtime.ReadMemStats(&m)
	fmt.Printf(details+"\n", elapsed, bToMb(m.Alloc), m.Alloc, m.NumGC)
	log.Printf(details, elapsed, bToMb(m.Alloc), m.Alloc, m.NumGC)
}
