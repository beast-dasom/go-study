package main

import "fmt"

func main() {

	a := [5]int{32, 29, 57, 16, 81}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println()

	for i, value := range a {
		fmt.Println(i, value)
	}

	for value := range a { // index만 표기됨
		fmt.Println(value)
	}

	for _, value := range a { // _, 를 붙이면 value 값 만 표기 됨.
		fmt.Println(value)
	}

}
