package main

import (
	"fmt"
)

func main() {
	fmt.Println("functions in golang")
	// greeter()
	// result := adder(5, 10)
	result := proAdder(2, 8, 9, 1, 2, 10, 55, 94, 945)
	fmt.Println("Result is: ", result)

}

// simple function
func greeter() {
	fmt.Println("Namstey from golang")
}

// parameter function with return
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// add n number of values
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
