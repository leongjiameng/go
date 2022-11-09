package main

import (
	"fmt"
	"time"
)

func main() {
	current_year := time.Now().Year()
	for born_year := 1997; born_year <= current_year; born_year++ {
		fmt.Println(born_year)
	}
}
