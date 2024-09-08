package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go Dowerk(done)
	time.Sleep(2 * time.Second)
	close(done)
	// time.Sleep(2 * time.Second)
}
func Dowerk(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Print("channel closed")
			return
		default:
			fmt.Println("working progressing")
		}
	}
}
