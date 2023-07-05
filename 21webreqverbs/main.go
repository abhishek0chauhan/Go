package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web verb in golang")
	// PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Content length is: ", res.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(res.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "htpp://localhost:8080/post"

	//fake json data

	requestBody := strings.NewReader(`
	{
		"coursename":"golang",
		"price": 0
	}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8080/postform"

	// formdata

	data := url.Values{}

	data.Add("firstname", "Abhishek")
	data.Add("lastname", "Chauhan")
	data.Add("email", "abc@gmail.com")

	res, err := http.Post(myurl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Boys)

	fmt.Println(string(content))
}
