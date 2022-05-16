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
	valCtx := context.WithValue(childCtx, "k1", "v1")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(valCtx, wg)
	wg.Wait()
	fmt.Println("main completed")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(2)
	go f1(ctx, wg)
	go f2(ctx, wg)
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

func f1(ctx context.Context, wg *sync.WaitGroup) {
	//f1Ctx, _ := context.WithCancel(ctx)
	/* defer func() {
		fmt.Println("cancel called from defer")
		cancel()
	}() */

	f1Ctx := context.WithValue(ctx, "ck1", "child value 1")
	wg.Add(2)
	go f11(f1Ctx, wg)
	go f12(f1Ctx, wg)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f1 done!")
			break LOOP
		default:
			fmt.Print("f1")
			time.Sleep(500 * time.Millisecond)
		}
	}
	wg.Done()
}

func f11(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("Data from context [key=k1]= ", ctx.Value("k1"))
	fmt.Println("Data from context [key=ck1]= ", ctx.Value("ck1"))
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f11 done!")
			break LOOP
		default:
			fmt.Print("f11")
			time.Sleep(200 * time.Millisecond)
		}
	}
	wg.Done()
}

func f12(ctx context.Context, wg *sync.WaitGroup) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f12 done!")
			break LOOP
		default:
			fmt.Print("f12")
			time.Sleep(300 * time.Millisecond)
		}
	}
	wg.Done()
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f2 done!")
			break LOOP
		default:
			fmt.Print("f2")
			time.Sleep(700 * time.Millisecond)
		}
	}
	wg.Done()
}
