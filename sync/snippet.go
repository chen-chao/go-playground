// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-07-25 13:40:15

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"fmt"
	"sync"
	"time"
)

type Quota struct {
	available int
	timelimit time.Duration
	expired   chan int
	backlog   chan int
	lock      sync.Mutex
	wg        sync.WaitGroup
	ok        chan struct{}
}

func (q *Quota) accept(size int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.available += size
}

func (q *Quota) release(size int) {
	q.accept(size)
}

func (q *Quota) consume(size int) {
	q.accept(-size)
}

func (q *Quota) Run() {
	var query int
	for {
		select {
		case size := <-q.expired:
			q.release(size)
			if query > 0 && q.available > query {
				query = 0
				q.ok <- struct{}{}
			}
		case query = <-q.backlog:
		}
	}
}

func (q *Quota) Stop() {
	q.wg.Wait()
	close(q.expired)
}

func (q *Quota) Wait(size int) {
	if q.available > size {
		q.consume(size)
		q.wg.Add(1)
		go func() {
			defer q.wg.Done()
			time.Sleep(q.timelimit)
			q.expired <- size
		}()
	} else {
		q.backlog <- size
		<-q.ok
		q.Wait(size)
	}
}

func NewQuota(limit int, timeout time.Duration) *Quota {
	expired := make(chan int)
	backlog := make(chan int)
	ok := make(chan struct{})

	return &Quota{
		available: limit,
		timelimit: timeout,
		expired:   expired,
		backlog:   backlog,
		ok:        ok,
	}
}

func main() {
	quota := NewQuota(10, 3*time.Second)
	go quota.Run()

	waiting := []int{2, 3, 4, 1, 6, 4, 7}
	start := time.Now()

	for _, w := range waiting {
		fmt.Println("Waiting", w)
		go func(w int) {
			fmt.Println("Before:\t", time.Since(start).Round(time.Second))
			quota.Wait(w)
			fmt.Println("After:\t", time.Since(start).Round(time.Second))
		}(w)
	}
	quota.Stop()
}
