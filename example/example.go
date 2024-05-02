package main

import (
	"fmt"
	"sync"

	"github.com/twistingmercury/queue"
)

func main() {
	q := queue.New[int]()

	// Enqueue items
	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Create multiple dequeuing goroutines
	numDequeuers := 3
	wg.Add(numDequeuers)

	for i := 0; i < numDequeuers; i++ {
		go func() {
			defer wg.Done()

			// Dequeue items until the queue is empty
			for {
				item, ok := q.Dequeue()
				if !ok {
					// Queue is empty, exit the goroutine
					break
				}
				fmt.Printf("Dequeued item: %d\n", item)
			}
		}()
	}

	// Wait for all dequeuing goroutines to finish
	wg.Wait()

	fmt.Println("All items dequeued")
}
