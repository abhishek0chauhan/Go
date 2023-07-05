package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("web requests in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is %T", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// convert data bytes into content
	content := string(databytes)
	fmt.Println(content)
}
