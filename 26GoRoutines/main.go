package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main(){
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist{
		go getStatusCode(web)
		// time.Sleep(3 * time.Second)
		wg.Add(1)
	}

	wg.Wait()
}


func getStatusCode(endpoint string){
	defer wg.Done()
	response, err := http.Get(endpoint)
	if err != nil{
		fmt.Println("Something when wrongggggg")
	}else{
		fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
	}
}