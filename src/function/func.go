package main

import "fmt"

func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func 한글함수() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

func sumAndDiff2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func variableness(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func main() {
	sum, diff := sumAndDiff(6, 2)
	fmt.Println(sum, diff)
	fmt.Println(한글함수())

	a, _, c, _, e := 한글함수()
	fmt.Println(a, c, e)

	fmt.Println(sumAndDiff2(6, 2))

	sum2, diff2 := sumAndDiff2(6, 2)
	fmt.Println(sum2, diff2)

	r := variableness(1, 2, 3, 4, 5)
	fmt.Println(r)

	n := []int{1, 2, 3, 4, 5}
	p := variableness(n...)

	fmt.Println(p)
}
