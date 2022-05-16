package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

//Communicate by sharing memory
type AtomicCounter struct {
	counter int
	sync.Mutex
}

func (ac *AtomicCounter) increment() {
	ac.Lock()
	{
		ac.counter++
	}
	ac.Unlock()
}

var atomicCounter AtomicCounter

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
	fmt.Printf("Counter = %d\n", atomicCounter.counter)
}

func f1(id int) {
	atomicCounter.increment()
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
