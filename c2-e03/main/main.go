package main

import "fmt"

func main(){
	const x int= 5
	const y=100

	const(
		a=5
		b int = 40
	)

	fmt.Println(x);
	fmt.Println(y);

	fmt.Println(a);
	fmt.Println(b);
}