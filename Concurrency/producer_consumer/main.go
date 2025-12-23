package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(jobs chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing job:", i)
		jobs <- i               // send job
		time.Sleep(500 * time.Millisecond)
	}
	close(jobs) // very important
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Consumer %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}
// Backpressure:
// If consumers process jobs slower than producers send them,
// the channel buffer eventually fills up.
// Once the buffer is full, producers block on send (jobs <- x).
// This automatically slows down producers and prevents
// unbounded memory usage or system overload.

func main() {
	jobs := make(chan int, 3) // buffered channel

	var wg sync.WaitGroup

	// start consumers
	numConsumers := 2
	wg.Add(numConsumers)

	for i := 1; i <= numConsumers; i++ {
		go consumer(i, jobs, &wg)
	}

	// start producer
	go producer(jobs)

	// wait for consumers to finish
	wg.Wait()
	fmt.Println("All jobs processed")
}