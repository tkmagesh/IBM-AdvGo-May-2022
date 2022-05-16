package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
	fmt.Println("Hit ENTER to exit....")
	var input string
	fmt.Scanln(&input)

}

func f1(id int) {
	rnd := rand.Intn(5000)
	time.Sleep(time.Duration(rnd * int(time.Millisecond)))
	/* if rnd%7 == 0 {
		fmt.Printf("Hit ENTER for f1[%d] to complete\n", id)
		var input string
		fmt.Scanln(&input)
	} */
	completionTime := time.Now()
	fmt.Printf("f1[%d] completed at %v\n", id, completionTime)

	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
