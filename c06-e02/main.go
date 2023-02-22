package main

import "fmt"

func main() {
	var nums = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x := foo(nums...)
	y := bar(nums)
	fmt.Println(x)
	fmt.Println(y)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

/*
Hands-on exercise #2
● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in
*/
