package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hvwjdhadavc"

func main() {
	fmt.Println("urls in golang")
	fmt.Println(myurl)

	//parsing url
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	// fmt.Printf("The type of query params are %T\n", qparams)

	// fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// don't forget the `&` when constructing url its pass reference not copy
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=abhishek",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
