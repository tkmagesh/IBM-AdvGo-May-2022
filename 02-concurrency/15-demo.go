package main

import "fmt"

func main() {
	fmt.Println("main started")
	count := 10
	ch := genPrimes(3, count)
	/* get the prime numbers and print them */
	for i := 0; i < count; i++ {
		fmt.Println("Prime No :", <-ch)
	}
	fmt.Println("main completed")
}

func genPrimes(start, count int) chan int {
	/* keep generating the "count" number of prime numbers starting from  "start" */
	ch := make(chan int)
	go func() {
		for count > 0 {
			if isPrime(start) {
				ch <- start
				count--
			}
			start++
		}
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
