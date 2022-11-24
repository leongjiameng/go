package main

import "fmt"

/*
Using a COMPOSITE LITERAL:
● create an ARRAY which holds 5 VALUES of TYPE int
● assign VALUES to each index position.
● Range over the array and print the values out.
● Using format printing
○ print out the TYPE of the array
*/
func main() {
	//var x [5]int
	//fmt.Println(x)
	//x[0] = 1
	//x[1] = 2
	//x[2] = 3
	//x[3] = 4
	//x[4] = 5
	x := [5]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)

}
