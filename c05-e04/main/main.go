package main

import "fmt"

func main() {
	p1 := struct {
		last              string
		first             string
		favouriteIceCream []string
	}{
		first: "James",
		last:  "Bond",
		favouriteIceCream: []string{
			"Vanilla",
			"Yam",
		},
	}

	p2 := struct {
		last              string
		first             string
		favouriteIceCream []string
	}{
		first: "Jackie",
		last:  "Chan",
		favouriteIceCream: []string{
			"Chocolate",
			"Mocha",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
