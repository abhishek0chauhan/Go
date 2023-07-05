package main

import "fmt"

func main() {
	fmt.Println("wlecome to the class of array")

	var fruitList [5]string

	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Tomato"
	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruit list is ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("veg list is ", vegList)
	fmt.Println("veg list is ", len(vegList))
}
