package main

import "fmt"

func main() {
	var num int = 1
	var numPtr *int = &num

	fmt.Println(numPtr)
	fmt.Println(&num) // &을 붙이면 메모리 주소를 구함.

}
