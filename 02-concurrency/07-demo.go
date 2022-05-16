package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

//Communicate by sharing memory
var counter int64

func main() {
	fmt.Println("main started")
	strCount := os.Args[1]
	count, err := strconv.Atoi(strCount)
	if err != nil {
		log.Fatalln("Invalid command line arguments. Usage => ./04-demo [count]")
	}
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go f1(i) //scheduling f1 to execute through the scheduler
	}
	f2()
	wg.Wait() //block until the counter becomes 0
	fmt.Println("main completed")
	fmt.Printf("Counter = %d\n", counter)
}

func f1(id int) {
	atomic.AddInt64(&counter, 1)
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
