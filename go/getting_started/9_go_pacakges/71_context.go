package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	//defer cancel()
	//ctx := context.TODO()
	go longProcess(ctx, ch)
	cancel()

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
            fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%")
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("################")
}
