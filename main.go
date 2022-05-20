package main

import "fmt"

func main() {
	// call power multiple times
	power(2, 8)
	power(10, 4)
	power(3, 7)
	power(123, 321)
	power(4, 0)

	// we can also use the returned value
	num := power(2, 8)
	fmt.Println("Look at this number:", num)
}

// cube multiplies an integer by itself 3 times.
func cube(x int) int {
	return x * x * x
}

// power multiplies x by itself y times.
func power(x int, y int) int {
	if y == 0 {
		fmt.Println(
			fmt.Sprintf("%d^%d = 1", x, y),
		)
		return 1
	}

	num := x
	for i := 1; i < y; i++ {
		num = num * x
	}

	// pretty print
	fmt.Println(
		fmt.Sprintf("%d^%d = %d", x, y, num),
	)

	return num
}
