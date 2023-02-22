package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	aTruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: false,
	}

	aSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "brown",
		},
		luxury: true,
	}

	fmt.Println(aTruck)
	fmt.Println(aSedan)
	fmt.Println(aTruck.doors)
	fmt.Println(aSedan.color)
}
