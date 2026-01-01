package examples

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int){
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func DemoWaitGroups(){
	var wg sync.WaitGroup
	numWorkers := 5

	for i:=1; i<=numWorkers; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()
	fmt.Println("All workers completed")
}