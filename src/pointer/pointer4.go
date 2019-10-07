package main

import "fmt"

func hello2(n *int) {
	*n = 2
}

func main() {
	var n int = 1

	hello2(&n)

	fmt.Println(n)
}
