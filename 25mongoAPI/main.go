package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhishek0chauhan/mongoAPI/router"
)

func main() {
	fmt.Println("APIs using mongodb")
	r := router.Router()
	fmt.Println("server is getting started....")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 4000...")
}
