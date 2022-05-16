package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
		time.Sleep(500 * time.Millisecond)
		ch <- 20
		time.Sleep(500 * time.Millisecond)
		ch <- 30
		time.Sleep(500 * time.Millisecond)
		ch <- 40
		time.Sleep(500 * time.Millisecond)
	}()
	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}
