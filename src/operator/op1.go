package main

import "fmt"

func main() {
	a := 4
	b := 3
	var c int
	c = 5
	var d = 6
	var e int = 3
	var f float32 = 3.14

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Println("result2 = ", a^b)
}
