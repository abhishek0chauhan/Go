package main

import "fmt"

func main() {
	fmt.Println("Strcucts in golang")

	// no inheritance in golang; No super or parent

	abhishek := User{"Abhishek", "abhi@gmail.com", true, 23}

	// fmt.Println("struct user: ", abhishek)
	// fmt.Printf("abhishek details are: %+v\n", abhishek)
	// fmt.Printf("abhishek name is %v and email is %v", abhishek.Name, abhishek.Email)

	abhishek.GetStatus()
	abhishek.NewMail()
	fmt.Printf("abhishek name is %v and email is %v", abhishek.Name, abhishek.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// struct pass in func make methods in golang
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "new@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
