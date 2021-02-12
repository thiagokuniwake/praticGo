package main

import (
	"fmt"
	"sync"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		waitGroup.Done()
	}()

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}
}
