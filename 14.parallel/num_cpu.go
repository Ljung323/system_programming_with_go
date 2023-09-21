package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Get the current MAXPROCS
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Println("Current MAXPROCS:", maxProcs)

	// Get the number of CPUs
	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numCPU)

	// Set MAXPROCS to the number of CPUs
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("New MAXPROCS:", numCPU)

	// Get the current MAXPROCS
	maxProcs = runtime.GOMAXPROCS(0)
	fmt.Println("Current MAXPROCS:", maxProcs)
}
