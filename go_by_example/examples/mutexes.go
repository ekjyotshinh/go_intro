package examples

import (
	"fmt"
	"sync"
)

type Container struct{
	mu   sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++	// Critical section
}
func DemoMutexes(){
	c := Container{
		counters: map[string]int{"a":0, "b":0},
	}
	var wg sync.WaitGroup
	numIncrements := 10000

	for i:=0; i<numIncrements; i++{
		wg.Add(3)
		go func() {
			defer wg.Done()
			c.inc("a")
		}()
		go func() {
			defer wg.Done()
			c.inc("a")
		}()
		go func() {
			defer wg.Done()
			c.inc("b")
		}()
	}

	wg.Wait()
	fmt.Printf("Final counters: %+v\n", c.counters)
}