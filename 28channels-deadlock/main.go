package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels and deadlock")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup){
		close(myCh)
		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 5
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}