package main

import (
	"fmt"
	"time"
)

func main() {
	current_year := time.Now().Year()
	if current_year == 2023 {
		fmt.Println("The year is 2023")
	} else if current_year == 1997 {
		fmt.Println("The year is 1997")
	} else {
		fmt.Println("Neither")
	}
}
