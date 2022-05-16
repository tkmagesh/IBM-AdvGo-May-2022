package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	mainCh := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		val := <-mainCh
		fmt.Println(val)
	}()
	go f1(ch1)
	go f2(ch2)
	for i := 0; i < 3; i++ {
		select {
		case mainCh <- 500:
			fmt.Println("Sent the data to the main channel")
		case value1 := <-ch1:
			fmt.Println(value1)
		case value2 := <-ch2:
			fmt.Println(value2)
			/* default:
			fmt.Println("No channel operations were successful") */
		}
	}
}

func f1(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 100
}

func f2(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 200
}
