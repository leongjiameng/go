package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("Done")
}

/*
Hands-on exercise #6
● Build and use an anonymous func
*/
