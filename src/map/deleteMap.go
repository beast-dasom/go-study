package main

import "fmt"

func main() {
	a := map[string]int{"hello": 10, "world": 20}

	delete(a, "world")

	fmt.Println(a)

}
