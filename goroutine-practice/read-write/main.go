package main

import (
	"fmt"
	"sync"
)

type handle struct {
	writewg sync.WaitGroup
	readwg  sync.WaitGroup
	mu      sync.RWMutex
}

func (h *handle) Writer(k int, ch chan<- int) {
	h.mu.RLock()
	defer h.writewg.Done()
	defer h.mu.RUnlock()
	for i := 0; i < 10; i++ {
		fmt.Printf("writer %d writes %d'th book \n", k, i)
		ch <- i
	}
}

func (h *handle) Reader(k int, ch <-chan int) {
	defer h.readwg.Done()
	for v := range ch {
		fmt.Printf("reader %d reads %d book \n", k, v)
	}
}

func main() {
	h := handle{}
	ch := make(chan int)

	writers := 2
	readers := 4

	// Start writers
	for i := 1; i <= writers; i++ {
		h.writewg.Add(1)
		go h.Writer(i, ch)
	}

	// Close channel after all writers are done
	// // Start readers
	go func() {
		h.writewg.Wait() // Wait for all writers to finish
		close(ch)        // Close the channel so readers can stop
	}()

	for i := 1; i <= readers; i++ {
		h.readwg.Add(1)
		go h.Reader(i, ch)
	}
	h.readwg.Wait() // Wait for all readers to finish
}
