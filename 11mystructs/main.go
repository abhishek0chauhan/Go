package main

import "fmt"

func main() {
	fmt.Println("Strcucts in golang")

	// no inheritance in golang; No super or parent

	abhishek := User{"Abhishek", "abhi@gmail.com", true, 23}

	fmt.Println("struct user: ", abhishek)
	fmt.Printf("abhishek details are: %+v\n", abhishek)
	fmt.Printf("abhishek name is %v and email is %v", abhishek.Name, abhishek.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
