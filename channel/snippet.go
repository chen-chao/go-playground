// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-06-25 12:53:19

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go produce(ch)
	consume(ch)
	fmt.Println("Main:  \t ", time.Now())

}

func produce(ch chan<- int) {
	for i := 0; ; i++ {
		fmt.Println("Produce:\t", time.Now(), i*i)
		ch <- i * i
		if i == 8 {
			close(ch)
			fmt.Println("Produce:\t", time.Now(), "Channel Closed")
			break
		}
	}
}

func consume(ch <-chan int) {
	var wg sync.WaitGroup
	for i := range ch {
		fmt.Println("Consume:\t", time.Now(), i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Goroutine:\t", time.Now(), i)
		}(i)
	}
	go func() {
		wg.Wait()
	}()
}
