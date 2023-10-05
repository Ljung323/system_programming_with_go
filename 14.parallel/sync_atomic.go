package main

import (
	"fmt"
	"sync/atomic"
)

var id int64

func generateId() int64 {
	return atomic.AddInt64(&id, 1)
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId())
		}()
	}
}
