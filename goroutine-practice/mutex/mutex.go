package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
	ch    chan int
}

func (c *Counter) Increment(i int) {
	c.mu.Lock()
	c.ch <- i
	c.mu.Unlock()
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return <-c.ch
}

func main() {
	counter := &Counter{ch: make(chan int)}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			counter.Increment(i)
		}()
	}
	for forval := range counter.ch {
		go counter.GetCount() // prints 10
		fmt.Println(forval)
	}
	close(counter.ch)
	wg.Wait()
	
}
