package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	year := 1997
	for year <= t.Year() {
		fmt.Println(year)
		year++
	}
}
