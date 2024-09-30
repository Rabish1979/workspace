package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	go performTask(ctx)
	time.Sleep(3 * time.Second)

	ctx2, cancel2 := context.WithDeadlineCause(context.Background(), time.Now().Add(2*time.Second), errors.New("RPC Timeout"))

	defer cancel2()
	go PerformTask2(ctx2)
	// Simulate work
	time.Sleep(3 * time.Second)

	// Print the error cause
	fmt.Println(ctx2.Err()) // prints "context deadline exceeded: RPC timeout"
}

func PerformTask2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task2 completed or deadline exceeded:", ctx.Err())
		return
	}
}

func performTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task completed or deadline exceeded:", ctx.Err())
		return
	}
}
