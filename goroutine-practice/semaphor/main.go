package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{} // Acquire semaphore
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d finished\n", id)
	<-sem // Release semaphore
}

func main() {
	sem := make(chan struct{}, 3) // Only allow 3 workers at a time
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait()
}
