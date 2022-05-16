package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(10)
	go f1(wg) //scheduling f1 to execute through the scheduler
	f2()
	wg.Wait() //block until the counter becomes 0
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
