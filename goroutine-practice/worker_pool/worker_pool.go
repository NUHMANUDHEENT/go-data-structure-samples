package main

import (
	"fmt"
	"strings"
	"sync"
)

// func main() {
// 	n := time.Now()
// 	job := 10
// 	workers := 3

// 	jobchan := make(chan int, job)
// 	result := make(chan int, job)
// 	var wg sync.WaitGroup

// 	// start worker goroutine
// 	for i := 1; i < workers; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			worker(i, jobchan, result)
// 		}(i)
// 	}
// 	for i := 1; i < job; i++ {
// 		jobchan <- i
// 	}
// 	close(jobchan)

//		//wait for all workers to finish
//		go func() {
//			wg.Wait()
//			close(result)
//		}()
//		// collect results from result channel
//		for v := range result {
//			fmt.Println("result is ", v)
//		}
//		fmt.Println("time is ", time.Since(n))
//	}
//
//	func worker(i int, job <-chan int, result chan<- int) {
//		for v := range job {
//			fmt.Println("worker processing", i, v)
//			result <- v * 2
//		}
//	}
func main() {
	str := "ds ds gg 44 33 fd gg 44 33 qw hht dsd as"
	var wg sync.WaitGroup
	arr := strings.Split(str, " ")
	worker := len(arr) / 4
	temp := make(map[string]int)
	chan1 := make(chan string, len(arr))
	for _, v := range arr {
		chan1 <- v
	}
	close(chan1)
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(&temp, chan1)
		}()
	}

	wg.Wait()

	for val, v := range temp {
		fmt.Printf("%v == %v \n", val, v)
	}
}
func Worker(temp *map[string]int, chan1 <-chan string) {
	for v := range chan1 {
		(*temp)[v]++
	}
}
