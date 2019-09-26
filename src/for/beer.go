package main

import "fmt"

func main() {
	for bottles := 99; 0 < bottles; bottles-- {

		if bottles > 1 {
			fmt.Println(bottles, " bottles of beer on the wall, ", bottles, " bottles of beer.")
			fmt.Println("Take one down, pass it around, ", bottles-1, "bottles of beer on the wall.")
		} else {
			fmt.Println(bottles, " bottles of beer on the wall, ", bottles, " bottles of beer.")
			fmt.Println("Take one down, pass it around, No more bottles of beer.")
		}
	}
	fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
	fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
}
