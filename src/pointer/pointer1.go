package main

import "fmt"

func main() {
	var numPtr *int
	fmt.Println(numPtr)

	var numPtr2 *int = new(int)
	fmt.Println(numPtr2)

	var numPtr3 *int = new(int)

	*numPtr3 = 1

	fmt.Println(*numPtr3)
}
