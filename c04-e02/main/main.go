package main

import "fmt"

/*
● Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign 10 VALUES
● Range over the slice and print the values out.
● Using format printing
○ print out the TYPE of the slice
*/
func main() {
	x := []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)

}
