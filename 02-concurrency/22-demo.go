package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	stop := make(chan bool)

	go func() {
		var input string
		fmt.Scanln(&input)
		stop <- true
	}()

	ch := genPrimes(3, stop)
	/* get the prime numbers and print them */
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}

	fmt.Println("main completed")
}

func genPrimes(start int, stop <-chan bool) <-chan int {
	/* keep generating the prime numbers starting from  "start" for the duration "d' */
	ch := make(chan int)
	no := start
	go func() {
	LOOP:
		for {
			if !isPrime(no) {
				no++
				select {
				case <-stop:
					break LOOP
				default:
					continue LOOP
				}
			}
			select {
			case ch <- no:
				time.Sleep(500 * time.Millisecond)
				no++
			case <-stop:
				break LOOP
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

//our own implementation of time.After()
func timeout(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}
