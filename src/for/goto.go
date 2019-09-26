package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		goto ERROR1
	}

	if a == 2 {
		goto ERROR2
	}

	fmt.Println(a)
	fmt.Println("Success")

ERROR1:
	fmt.Println("Error 1")
	return

ERROR2:
	fmt.Println("Error 1")
	return
}
