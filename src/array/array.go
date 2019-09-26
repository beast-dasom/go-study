package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	a[2] = 7

	b := [5]int{
		1,
		2,
		3,
		4,
		5, // 여러줄로 나열 할 경우 마지막에 comma 붙임.
	}

	c := [...]string{"Ahri", "Fizz", "Zed"}

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)
}
