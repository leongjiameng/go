package main

import "fmt"

func main() {
	xi := []int{4, 5, 6, 7, 8, 9, 10}
	oddList := odd(multiply, xi...)
	fmt.Println(oddList)
}

func multiply(xi ...int) int {
	total := 1
	for _, v := range xi {
		total *= v
	}
	return total
}

func odd(f func(xi ...int) int, vi ...int) int {
	var oddList []int
	for _, v := range vi {
		if v%2 != 0 {
			oddList = append(oddList, v)
		}
	}
	return f(oddList...)
}

/*
Hands-on exercise #8
● Create a func which returns a func
● assign the returned func to a variable
● call the returned func
*/
