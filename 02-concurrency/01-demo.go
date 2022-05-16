package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() //scheduling f1 to execute through the scheduler
	f2()
	//DO NOT DO THIS
	//time.Sleep(2 * time.Second)

	//DO NOT DO THIS
	var input string
	fmt.Scanln(&input)
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
