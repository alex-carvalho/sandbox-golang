package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	Job    Job
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(time.Second) // Simulate work
		
		result := Result{
			Job:    job,
			Output: fmt.Sprintf("Processed by worker %d: %s", id, job.Data),
		}
		results <- result
	}
}

func workerPool() {
	const numWorkers = 3
	const numJobs = 5
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("task-%d", j)}
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Result: %s\n", result.Output)
	}
}

func pipeline() {
	input := make(chan int, 10)
	output := make(chan int, 10)
	
	// Stage 1: multiply by 2
	go func() {
		defer close(output)
		for n := range input {
			output <- n * 2
		}
	}()
	
	// Send data
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()
	
	// Receive results
	for result := range output {
		fmt.Printf("Pipeline result: %d\n", result)
	}
}

func fanOut() {
	input := make(chan int)
	output1 := make(chan int)
	output2 := make(chan int)
	
	// Fan-out: distribute work
	go func() {
		defer close(output1)
		defer close(output2)
		for i := range input {
			if i%2 == 0 {
				output1 <- i
			} else {
				output2 <- i
			}
		}
	}()
	
	var wg sync.WaitGroup
	
	// Worker 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range output1 {
			fmt.Printf("Even worker: %d\n", n)
		}
	}()
	
	// Worker 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range output2 {
			fmt.Printf("Odd worker: %d\n", n)
		}
	}()
	
	// Send data
	go func() {
		defer close(input)
		for i := 1; i <= 6; i++ {
			input <- i
		}
	}()
	
	wg.Wait()
}

func main() {
	fmt.Println("=== Worker Pool ===")
	workerPool()
	
	fmt.Println("\n=== Pipeline ===")
	pipeline()
	
	fmt.Println("\n=== Fan-Out ===")
	fanOut()
}