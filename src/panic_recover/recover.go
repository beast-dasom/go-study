package main

import "fmt"

func f() {
	defer func() {
		s := recover()

		fmt.Println(s)
	}()

	panic("Error !!")
}

func main() {
	f()
	fmt.Println("Hello, world")
}
