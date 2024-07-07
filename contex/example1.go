package main

import (
	"context"
	"fmt"
	"runtime"
)

func doWork(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			// If the context is canceled, it will be executed here
			return
		default:
			// Work logic is placed here
		}
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	doWork(ctx)
	// Some code...
	cancel() // This causes the goroutines waiting for ctx.Done() to finish

	defer fmt.Println(runtime.NumGoroutine())
}
