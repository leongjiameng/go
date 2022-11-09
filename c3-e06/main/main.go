package main

import (
	"fmt"
	"time"
)

func main() {
	current_year := time.Now().Year()
	if current_year == 2022 {
		fmt.Println("The year is 2022")
	} else {
		fmt.Println("The year is not 2022")
	}
}
