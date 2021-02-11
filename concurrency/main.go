package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func init() {
	//Force the qty of CPU exec. in parallel
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	waitGroup.Add(2) //Qty. process to wait
	go runProcess("P1", 10)
	go runProcess("P2", 10)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}

	waitGroup.Done()
}
