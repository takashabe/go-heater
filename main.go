package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(
		sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	worker := runtime.NumCPU
	for i := 0; i < worker; i++ {
		go heating(ctx)
	}

	select {
	case <-sigCh:
		fmt.Println("receive signal...")
		return
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	}
}

func heating(ctx context.Context) {
	var i int64
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if i >= math.MaxInt64 {
				i = 0
			}
			i++
		}
	}
}
