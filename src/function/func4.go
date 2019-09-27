package main

import "fmt"

func sum2(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := []func(int, int) int{sum2, diff}

	fmt.Println(f[0](1, 2))
	fmt.Println(f[1](1, 2))

	m := map[string]func(int, int) int{
		"sum":  sum2,
		"diff": diff,
	}

	fmt.Println(m["sum"](1, 2))
	fmt.Println(m["diff"](1, 2))
}
