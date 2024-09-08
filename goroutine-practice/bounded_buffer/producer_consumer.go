package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, done chan bool) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
	done <- true
}

func consumer(ch chan int, done chan bool) {
	for value := range ch {
		fmt.Printf("Consumed: %d\n", value)
		time.Sleep(time.Millisecond * 1500)
	}
	done <- true
}

func main() {
	ch := make(chan int) // Bounded buffer with size 3
	done := make(chan bool,3)

	go producer(ch, done)
	go consumer(ch, done)

	<-done
	<-done
}
