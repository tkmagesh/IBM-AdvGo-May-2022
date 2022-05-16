package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	childCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(childCtx, wg)
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
