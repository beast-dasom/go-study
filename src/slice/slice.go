package main

import "fmt"

func main() {

	a := make([]int, 5, 10)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := []int{1, 2, 3}

	b = append(b, 4, 5, 6)

	c := []int{7, 8, 9}

	b = append(b, c...)

	fmt.Println(b)
}
