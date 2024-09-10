package main

import (
	"context"
	"fmt"
	"time"
)

func demo1() {
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
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func main() {

}

func demo3() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	slowFunc := func(ctx context.Context, i int) {
		childCtx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
		defer cancel()

		fmt.Printf("time: %v query No. %d\n", time.Now(), i)
		select {
		case <-childCtx.Done():
			fmt.Printf("time: %v child context err: %v\n", time.Now(), childCtx.Err())
		}
	}

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("time: %v parent context err: %v\n", time.Now(), ctx.Err())
			return
		default:
			slowFunc(ctx, i)
		}
	}
}
func demo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
