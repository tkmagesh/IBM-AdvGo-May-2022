package main

import (
	"fmt"
	"sync"
)

func main() {
	//share memory by communicating
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
