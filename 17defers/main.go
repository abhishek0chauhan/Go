package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// fmt.Println("Defer in golang")

	// defer fmt.Println("Hello")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Defer in golang")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
