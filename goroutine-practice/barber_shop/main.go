package main

import (
	"fmt"
	"time"
)

const capacity = 5

func customer(id int, waiting chan struct{}) {
	select {
	case waiting <- struct{}{}:
		fmt.Printf("Customer %d is waiting.\n", id)
		time.Sleep(time.Second * 2)
		<-waiting
		fmt.Printf("Customer %d is getting a haircut.\n", id)
	default:
		fmt.Printf("Customer %d left (no space).\n", id)
	}
}

func main() {
	waiting := make(chan struct{}, capacity)

	for i := 1; i <= 10; i++ {
		go customer(i, waiting)
		time.Sleep(time.Millisecond * 300)
	}

	time.Sleep(time.Second * 10)
}
