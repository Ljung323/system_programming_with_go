package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 1:", i)
			time.Sleep(time.Millisecond * 100)
			runtime.Gosched() // Yield the CPU to other goroutines
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 2:", i)
			time.Sleep(time.Millisecond * 100)
			runtime.Gosched() // Yield the CPU to other goroutines
		}
	}()

	// Wait for the goroutines to complete
	fmt.Println("Waiting for goroutines to complete...")
	time.Sleep(time.Second * 3)
}

// expected output from copilot
// goroutine 1: 0
// goroutine 2: 0
// goroutine 1: 1
// goroutine 2: 1
// goroutine 1: 2
// goroutine 2: 2
// goroutine 1: 3
// goroutine 2: 3
// goroutine 1: 4
// goroutine 2: 4
// Waiting for goroutines to complete...
