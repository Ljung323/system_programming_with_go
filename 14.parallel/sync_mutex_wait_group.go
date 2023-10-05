package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	loop_count := 100
	wg.Add(loop_count)

	for i := 0; i < loop_count; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("done")
}
