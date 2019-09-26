package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	var b [3]int

	b = a
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3}
	var d []int

	d = c
	d[0] = 9

	fmt.Println(c)
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5}
	f := make([]int, 3)

	copy(f, e)

	fmt.Println(e)
	fmt.Println(f)

	f[0] = 9

	fmt.Println(f)

	z := []int{1, 2, 3, 4, 5}
	fmt.Println(z)
	fmt.Println(z[2:3])
	fmt.Println(z[3:5])

	fmt.Println(z[0:len(z)])
	fmt.Println(z[:4])
	fmt.Println(z[2:])

}
