package main

import "fmt"

func main() {
	a := func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}
	a()
	fmt.Println("Done")
}

/*
Hands-on exercise #7
● Assign a func to a variable, then call that func
*/
