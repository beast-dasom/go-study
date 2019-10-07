package main

import "fmt"

func HellowWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("hello")
	}()
}

func main() {
	HellowWorld()
}
