package main

import "fmt"

func main() {
	y := make([]string, 12, 12)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	states := []string{`Perlis`, `Kelantan`, `Kedah`, `Pahang`, `Penang`, `Perak`, `Melaka`, `Selangor`, `Sarawak`, `Sabah`, `Wilayah Persekutuan`, `Negeri Sembilan`}

	for i, v := range states {
		y[i] = v
	}

	fmt.Println(len(y))
	fmt.Println(cap(y))

	for j, k := range y {
		fmt.Println(j, k)
	}

	y = append(y, `Johor`)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}
