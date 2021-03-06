package main

import "fmt"

func main() {
	fmt.Println("main started")
	ch := make(chan int)
	count := 10
	go genPrimes(3, count, ch)
	/* get the prime numbers and print them */
	for i := 0; i < count; i++ {
		fmt.Println("Prime No :", <-ch)
	}
	fmt.Println("main completed")
}

func genPrimes(start, count int, ch chan int) {
	/* keep generating the "count" number of prime numbers starting from  "start" */
	for count > 0 {
		if isPrime(start) {
			ch <- start
			count--
		}
		start++
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
