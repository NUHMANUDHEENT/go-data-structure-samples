package main

import (
	"fmt"
	"sync"
)

// func PrintNumbers(n int, wg *sync.WaitGroup, ch chan int) {
// 	defer wg.Done()
// 	ch <- n
// }

func main() {
	n := 20
	producer := make(chan int)
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		cunsumer(producer)
	}()
	for i := 0; i < n; i++ {
		producer <- i
	}
	go func() {	
		wg.Wait()
	}()
}	
func cunsumer(producer chan int) {
	defer close(producer)
	// out := make(chan int)
	for v := range producer {
		fmt.Println(v * v)
	}
}
