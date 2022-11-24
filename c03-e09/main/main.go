package main

import (
	"fmt"
)

func main() {
	var favSport string = "Jogging"
	switch favSport {
	case "Badminton":
		fmt.Println("Badminton is the best")
	case "Swimming":
		fmt.Println("Swim like a fish")
	case "Jogging":
		fmt.Println("Run like a bunny")
	}
}
