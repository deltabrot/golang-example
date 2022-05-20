package main

import "fmt"

func jamie() {
	// variables

	// long hand
	var num int // == 0
	num = square(8)

	// short hand
	num2 := square(8)

	fmt.Println(num + 10)
	fmt.Println(num2)
}

func square(x int) int {
	return x * x
}

// for loop example
func loopMe(numberOfLoops int) {
	for i := 0; i < numberOfLoops; i++ {
		fmt.Println(i)
	}
}

// add adds two numbers together.
func add(x int, y int) int {
	return x + y
}
