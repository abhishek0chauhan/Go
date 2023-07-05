package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Vlaue of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Go one step ahead")
	case 2:
		fmt.Println("Go two step ahead")
	case 3:
		fmt.Println("Go three step ahead")
	case 4:
		fmt.Println("Go four step ahead")
	case 5:
		fmt.Println("Go five step ahead")
		fallthrough // case 5 and 6 are gona work
	case 6:
		fmt.Println("Go six step ahead and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}
