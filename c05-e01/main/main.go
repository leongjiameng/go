package main

import "fmt"

type person struct {
	last              string
	first             string
	favouriteIceCream string
}

func main() {
	p1 := person{
		first:             "James",
		last:              "Bond",
		favouriteIceCream: "Vanilla",
	}

	p2 := person{
		first:             "Jackie",
		last:              "Chan",
		favouriteIceCream: "Chocolate",
	}

	fmt.Println(p1)
	fmt.Println(p2)

}
