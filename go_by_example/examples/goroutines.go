package examples

import (
	"fmt"
	"time"
)

func genericFunction(id int) {
	fmt.Printf("Goroutine %d is starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Goroutine %d is done\n", id)
}

func DemoGoroutines() {
	for i := 1; i <= 5; i++ {
		go genericFunction(i)	// Start a goroutine
	}

	// Wait for goroutines to finish
	// In a real application, use sync.WaitGroup or other synchronization methods
	time.Sleep(2 * time.Second)
	fmt.Println("All goroutines have completed")
}