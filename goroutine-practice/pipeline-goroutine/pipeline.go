package main

import "fmt"

func main() {
	values := []int{2, 4, 7, 3, 8, 12}
	// stage 1
	chanValues := chanData(values)
	// stage 2
	chanSquares := chanSquare(chanValues)
	// stage 3
	for val := range chanSquares {
		fmt.Println(val)
	}
}
func chanData(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}
func chanSquare(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
