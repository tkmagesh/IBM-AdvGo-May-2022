package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	childCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop")
	go fn(childCtx)
	var input string
	fmt.Scanln(&input)
	cancel()
	fmt.Println("main completed")
}

func fn(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fn done!")
			break LOOP
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}

}
