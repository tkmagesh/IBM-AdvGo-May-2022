package main

import "fmt"

func main() {
	/* 1. anonymous functions */
	func() {
		fmt.Println("Anonymous function invoked")
	}()

	func(x, y int) {
		fmt.Println("Add result = ", x+y)
	}(100, 200)

	result := func(x, y int) int {
		return x - y
	}(100, 200)
	fmt.Println("Subtract result =", result)

	/*
		quotient, remainder := func(x, y int) (int, int) {
			return x / y, x % y
		}(100, 7)
		fmt.Printf("Quotient = %d, Remainder = %d\n", quotient, remainder)
	*/
	quotient, remainder := func(x, y int) (q int, r int) {
		q = x / y
		r = x % y
		return
	}(100, 7)
	fmt.Printf("Quotient = %d, Remainder = %d\n", quotient, remainder)
	/*  */

	//var x interface{}
	var x any
	//x = 10

	x = "abc"
	//	x = true
	//	x = 3.14

	if val, ok := x.(int); ok {
		fmt.Println(val + 20)
	} else {
		fmt.Println("x is not an integer")
	}

	/* assigning functions to variables */
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var divide func(int, int) (int, int)
	divide = func(x, y int) (int, int) {
		return x / y, x % y
	}
	q, r := divide(100, 8)
	fmt.Println(q, r)

	//pass functions as arguments to other functions
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)

	//functions returned as return values

	addWithLog := getLogOperation(add)
	addWithLog(100, 200)

	subtractWithLog := getLogOperation(subtract)
	subtractWithLog(100, 200)
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		oper(x, y)
		fmt.Println("invocation completed")
	}
}

func logOperation(oper func(int, int), x, y int) {
	fmt.Println("invocation started")
	oper(x, y)
	fmt.Println("invocation completed")
}

/* func logAdd(x, y int) {
	fmt.Println("invocation started")
	add(x, y)
	fmt.Println("invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("invocation started")
	subtract(x, y)
	fmt.Println("invocation completed")
} */

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
