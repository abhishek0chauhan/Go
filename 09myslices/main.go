package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Welcome to the class of slices")

	// var fruitList = []string{"Apple", "Tomato"}

	// fmt.Printf("Type of the fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")

	// fmt.Println("fruitList ", fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println("fruitList ", fruitList)

	// fruitList = append(fruitList[1:2])
	// fmt.Println("fruitList ", fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println("fruitList ", fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println("fruitList ", fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 212
	// highScores[1] = 247
	// highScores[2] = 432
	// highScores[3] = 965

	// highScores[4] = 541

	// highScores = append(highScores, 415, 657, 741)

	// fmt.Println("highScores", highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// sort.Ints(highScores)

	// fmt.Println("sorted highScore", highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
