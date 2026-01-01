package examples

import (
	"fmt"
	"time"
)

// timoutes in Go can be handled using the time package along with goroutines and channels
// timeouts are useful to prevent operations from blocking indefinitely

func DemoTimeouts(){
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // simulate a long operation
		channel <- "Operation Completed"
	}()
	// in this case the operation will take longer than the timeout duration
	// so it would timeout and we would go to the timeout case
	select {
	case res := <- channel:
		fmt.Println(res)
	case <- time.After(1 *time.Second):
		fmt.Println("Timeout: Operation took too long")
	}

	channel2 := make(chan string)
	
	go func() {
		time.Sleep(500 * time.Millisecond) // simulate a shorter operation
		channel2 <- "Quick Operation Completed"
	}()

	// this time, the operation would complete before the timeout
	select {
	case res := <- channel2:
		fmt.Println(res)
	case <- time.After(1 *time.Second):
		fmt.Println("Timeout: Operation took too long")
	}

}