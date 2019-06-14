package builtin

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// examples are from go doc
func TestWithCancel(t *testing.T) {
	// gen generates int in a separate goroutine and sends them to the returned channel.
	// The callers of gen need to cancel the context once they are done consuming
	// the generated integers not to leak the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestWithDeadlin(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// even though ctx will be expired, it is good practice to call
	// its cancelation function in any case.
	// Failure to do so may keep the context and its parent alive longer than necessary
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("err: ", ctx.Err())
	}
}

func TestWithTimeOut(t *testing.T) {
	// same as context.WithDeadline(ctx, time.Now().Add(timeout))
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("err: ", ctx.Err())
	}
}
