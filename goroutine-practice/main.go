// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	const n = 3
// 	fibN := fib(n) // slow
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // producers sends the value to the channel
// func producers(id int, ch chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i:=1; i <= 100; i++ {
// 		fmt.Printf("goroutine %d sent the value %d\n", id, i)
// 		ch <- i
// 	}
// }

// // consumers consumes the values
// func consumers(id int, ch <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for val := range ch {
// 		fmt.Printf("value received by goroutine %d: %d\n",id, val)
// 	}
// }

// func main() {
// 	ch := make(chan int)
// 	var prowg sync.WaitGroup // separate waitgroup for producers
// 	var conswg sync.WaitGroup // separate waitgroup for consumers
	
// 	numProducers := 2
// 	numConsumers := 3

// 	for i:=1; i <= numProducers; i++ {
// 		prowg.Add(1)
// 		go producers(i, ch, &prowg)
// 	}

// 	for i:=1; i <= numConsumers; i++ {
// 		conswg.Add(1)
// 		go consumers(i, ch, &conswg)
// 	}

// 	go func ()  {
// 		prowg.Wait()
// 		close(ch)	
// 	}()
	
// 	conswg.Wait()
// }
package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id       int
	leftFork *sync.Mutex
	rightFork *sync.Mutex
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		p.think()
		p.eat()
	}
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking.\n", p.id)
	time.Sleep(time.Second)
}

func (p *Philosopher) eat() {
	// Lock both forks
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("Philosopher %d is eating.\n", p.id)
	time.Sleep(time.Second)

	// Unlock both forks
	p.rightFork.Unlock()
	p.leftFork.Unlock()
}

func main() {
	var forks = make([]*sync.Mutex, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &sync.Mutex{}
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:       i + 1,
			leftFork: forks[i],
			rightFork: forks[(i+1)%5], // Circular seating
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
}
