package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	childCtx, cancel := context.WithCancel(rootCtx)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(childCtx, wg)
	go func() {
		fmt.Println("Hit ENTER to stop")
		var input string
		fmt.Scanln(&input)
		cancel()
	}()
	wg.Wait()
	fmt.Println("main completed")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
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
	wg.Done()
}
