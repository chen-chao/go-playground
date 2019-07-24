// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-07-21 23:38:50

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"fmt"
	"sync"
	"time"
)

var tokens = make(chan struct{}, 3)

type foo struct {
	ch  chan int
	res chan int
	wg  sync.WaitGroup
}

var now = time.Now()

func (f *foo) Put(s int) {
	fmt.Println("Put:\t", time.Since(now).Round(time.Second), s)
	f.ch <- s
}

func (f *foo) Run() {
	for i := range f.ch {
		f.wg.Add(1)
		fmt.Println("WGADD:\t", f.wg)
		// tokens <- struct{}{}

		// Put tokens here would cause a deadlock.

		// if tokens is a non-buffered channel, we will never
		// reach the goroutine below.

		// if tokens is buffered channel, the for loop in far
		// quicker than the go func below. there will be a
		// time all tokens in the goroutines are released and
		// the tokens channel are filled with tokens, no
		// goroutine will be created to consume these tokens!
		// A deadlock!

		go func(i int) {
			tokens <- struct{}{}

			fmt.Println("Run:\t", time.Since(now).Round(time.Second), i)

			defer func() {
				f.wg.Done()
				fmt.Println("WGDONE:\t", f.wg)
			}()
			f.res <- i * i
			time.Sleep(5 * time.Second)
			<-tokens
		}(i)
	}
}

func (f *foo) Close() {
	close(f.ch)
	f.wg.Wait()
	close(f.res)
}

func (f *foo) Results() <-chan int {
	return f.res
}

func main() {
	ch := make(chan int)
	res := make(chan int)
	f := foo{ch: ch, res: res}

	go f.Run()

	for i := 0; i < 10; i++ {
		f.Put(i)
	}

	go f.Close()

	for r := range f.Results() {
		fmt.Println("Result:\t", time.Since(now).Round(time.Second), r)
	}

	fmt.Println("finished")
}
