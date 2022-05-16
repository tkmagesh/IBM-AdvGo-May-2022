package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hi there!"
	}()
	str := <-ch
	fmt.Println(str)
}
