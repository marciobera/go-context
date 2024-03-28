package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	dataChain := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChain <- result
			}()
		}
		wg.Wait()
		close(dataChain)
	}()

	for n := range dataChain {
		fmt.Printf("n = %d\n", n)
	}
}
