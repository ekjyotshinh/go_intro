package examples

import (
	"fmt"
	"time"
)

func DemoTimers(){

	timer1 := time.NewTimer(2 * time.Second)
	<- timer1.C // blocks until the timer1 channel sends a value
	fmt.Println("Timer 1 expired")

	// stopping a timer before it expires
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped before expiration")
	}

	// using time.After for simpler timer
	<- time.After(2 * time.Second)
	fmt.Println("Timer using time.After expired")
}