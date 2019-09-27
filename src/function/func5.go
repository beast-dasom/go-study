package main

import "fmt"

func main() {
	func() {
		fmt.Println("익명함수")
	}()

	func(s string) {
		fmt.Println(s)
	}("Hello World")

	r := func(a int, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(r)
}
