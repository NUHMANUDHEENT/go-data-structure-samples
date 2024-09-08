package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{3, 5, 6, 2, 7, 3}
	input := make(chan int, len(nums))

	var wg sync.WaitGroup
	for _, v := range nums {
		input <- v
	}
	close(input)
	workers := 3
	result := make(chan int)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for v := range input {
				fmt.Println(i)
				result <- v * 2
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println("result is ", v)
	}
}
