package main

import (
	"fmt"
)

//09-demo.go but without waitGroup

func main() {
	//share memory by communicating
	ch := make(chan int)
	fmt.Println("main started")

	go add(100, 200, ch)
	result := <-ch

	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
