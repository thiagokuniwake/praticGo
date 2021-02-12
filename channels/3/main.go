package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	// channel <- 10 - Deadlock!

	go func() {
		channel <- 10 //Channel has to be inside in a goroutines
	}()

	fmt.Println(<-channel)
}
