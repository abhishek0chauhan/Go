package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"sun", "tue", "wed", "fri", "sat"}

	fmt.Println(days)

	// for loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(i)
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("Value is %v\n", day)
	// }

	roughValue := 1

	// for roughValue < 11 {

	// 	if roughValue == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Value is: ", roughValue)
	// 	roughValue++
	// }

	for roughValue < 11 {

		// if roughValue == 5 {
		// 	roughValue++
		// 	continue
		// }

		if roughValue == 2 {
			goto freeocdecamp
		}

		fmt.Println("Value is: ", roughValue)
		roughValue++
	}

freeocdecamp:
	fmt.Println("Jumping at freeocdecamp.com")
}
