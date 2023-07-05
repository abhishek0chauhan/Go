package main

import "fmt"

func main() {
	fmt.Println("welcome to the class of maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "GoLang"

	fmt.Println("languages", languages)
	fmt.Println("JS short for: ", languages["JS"])

	delete(languages, "PY")

	fmt.Println("Languages", languages)

	// lops are interseting on golang

	for key, value := range languages {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}

	for _, value := range languages {
		// fmt.Println("key", key)
		fmt.Println("value", value)
	}

}
