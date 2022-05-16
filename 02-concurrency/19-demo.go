package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")

	ch := genPrimes(3, 5*time.Second)
	/* get the prime numbers and print them */
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}

	fmt.Println("main completed")
}

func genPrimes(start int, d time.Duration) <-chan int {
	/* keep generating the prime numbers starting from  "start" for the duration "d' */
	ch := make(chan int)
	go func() {
		no := start
		for {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				ch <- no
			}
			no++
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
