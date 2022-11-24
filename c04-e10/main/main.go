package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martini`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being Evil`, `Ice Cream`, `Sunsets`},
	}
	for i, v := range m {
		fmt.Printf("The %v stored value is\n", i)
		for j, k := range v {
			fmt.Printf("\tThe value no %v is %v\n", j, k)
		}
	}
	fmt.Println("After Delete")
	delete(m, "no_dr")
	for i, v := range m {
		fmt.Printf("The %v stored value is\n", i)
		for j, k := range v {
			fmt.Printf("\tThe value no %v is %v\n", j, k)
		}
	}
}
