package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hello, James"}}

	for i, v := range x {
		fmt.Printf("Record of %v \n", i)
		for j, k := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, k)
		}
	}
}
