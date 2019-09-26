package main

import "fmt"

func main() {
	a := 3 < 4 && 2 < 5
	fmt.Println(a)

	b := 3
	if b == 4 || b == 3 {
		fmt.Println("b는 3 또는 4 이다")
	} else {
		fmt.Println("b는 ", b)
	}

}
