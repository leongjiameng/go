package main

import (
	"fmt"
	"sync"
)

// Go Routine Wait Group
func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("About to exit")
}
