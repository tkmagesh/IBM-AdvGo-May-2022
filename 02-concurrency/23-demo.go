package main

import (
	"fmt"
)

var ch chan int = make(chan int)

func main() {
	fmt.Println("main started")
	go fn()
	/*
		fmt.Println("[@main] initiate the receive after 3 seconds")
		time.Sleep(3 * time.Second)
	*/
	fmt.Println("[@main] initiate the receive operation")

	data := <-ch
	fmt.Println("[@main], data = ", data)
	fmt.Println("main completed")
}

func fn() {
	fmt.Println("[@fn] attempting to send the data")
	ch <- 100
	fmt.Println("[@fn] attempt to send the data successful")
}

/*
	if the receive is already intiated (gate is open)
		then the send operation is a NON-Blocking operation

	if the receive is NOT intiated (gate is closed)
		then the send operation is a Blocking operation

	if the send is already intiated (gate is closed and the data is at the gate)
		then the receive operation is a NON-Blocking operation

	if the send is not intiated (gate is closed and the data is not at the gate)
		then the receive operation is a Blocking operation (opening the gate)

*/
