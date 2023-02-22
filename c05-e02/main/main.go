package main

import "fmt"

type person struct {
	last              string
	first             string
	favouriteIceCream []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favouriteIceCream: []string{
			"Vanilla",
			"Yam",
		},
	}

	p2 := person{
		first: "Jackie",
		last:  "Chan",
		favouriteIceCream: []string{
			"Chocolate",
			"Mocha",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favouriteIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("------------------")
	}
}
