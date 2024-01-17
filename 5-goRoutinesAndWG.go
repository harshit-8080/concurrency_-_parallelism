package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate some time-consuming task
func performTask(id int) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(time.Second * 2) // Simulating some time-consuming task
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	// Using a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launching multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Incrementing the WaitGroup counter

		go func(taskID int) {
			defer wg.Done() // Decrementing the WaitGroup counter when the goroutine completes
			performTask(taskID)
		}(i)
	}

	// Waiting for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed")
}
