package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")

	ch := genPrimes(3, 100)
	/* get the prime numbers and print them */
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}

	fmt.Println("main completed")
}

func genPrimes(start, end int) <-chan int {
	/* keep generating the "count" number of prime numbers starting from  "start" */
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				ch <- no
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
