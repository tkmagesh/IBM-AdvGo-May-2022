package main

import "fmt"

func main() {
	fmt.Println("main started")
	go genPrimes(3, 50)
	/* get the prime numbers and print them */
	fmt.Println("main completed")
}

func genPrimes(start, count int) {
	/* keep generating the "count" number of prime numbers starting from  "start" */
}
