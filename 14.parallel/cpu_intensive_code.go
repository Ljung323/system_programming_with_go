package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	// runtime.LockOSThread()
	// defer runtime.UnlockOSThread()

	// // SET MAXPROCS to 1
	// runtime.GOMAXPROCS(1)
	// procs := runtime.GOMAXPROCS(0)
	// fmt.Println("Current MAXPROCS:", procs)

	// Set the number of integers to sum
	n := 1000000000

	done := make(chan bool)
	start := time.Now()

	go func() {
		// Calculate the sum of the first n positive integers
		sum := big.NewInt(0)
		for i := 1; i <= n; i++ {
			sum = sum.Add(sum, big.NewInt(int64(i)))
		}

		done <- true
	}()

	// Wait for the goroutines to complete and combine the results
	<-done
	elapsed := time.Since(start)

	// Print the elapsed time
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
