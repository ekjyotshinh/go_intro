package examples

import (
	"fmt"
	"time"
)

// worker function that processes jobs from the jobs channel and sends results to the results channel
func worker(id int, jobs <- chan int, results chan <- int){
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // simulate time-consuming work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func DemoWorkerPools(){
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send jobs and close the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", <-results)
	}
}