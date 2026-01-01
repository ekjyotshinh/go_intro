package examples

import (
	"fmt"
	"time"
)

// tickers are for when you want to do something repeatedly at regular intervals.
func DemoTickers(){
	// tickers are similar to timers and use the same underlying mechanism but instead of firing once, they fire repeatedly at specified intervals.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func(){
		time.Sleep(1600 * time.Millisecond)
		done <- true
	}()

	for {
		select {
		case <- done:
			ticker.Stop()
			fmt.Println("Ticker stopped")
			return
		case t := <- ticker.C:
			fmt.Println("Ticker ticked at", t)
		}
	}
}