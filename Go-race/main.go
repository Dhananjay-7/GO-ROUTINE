package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex       // Create a mutex for synchronization
	var wg sync.WaitGroup   // Create a WaitGroup

	// Start multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done() // Decrement the WaitGroup counter when done

			for j := 0; j < 1000; j++ {
				mu.Lock()   // Lock the critical section
				counter++   // Safely increment the counter
				mu.Unlock() // Unlock after incrementing
			}
		}(&wg, &mu)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter:", counter)
}
