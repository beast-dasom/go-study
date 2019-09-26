package main

import "fmt"

func gugudan() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for j := 0; j < 11; j++ {
	// 	fmt.Println(j)
	// }

	for dan := 2; dan <= 9; dan++ {
		fmt.Printf("%d단\n", dan)

		if dan == 5 {
			continue
		}
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}
	}
}
