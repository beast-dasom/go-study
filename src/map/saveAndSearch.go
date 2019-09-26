package main

import "fmt"

func main() {

	solarSystem := make(map[string]float32)

	solarSystem["Mercury"] = 87.912
	solarSystem["Venus"] = 224.1581
	solarSystem["Earth"] = 381.01756

	fmt.Println(solarSystem["Earth"])

	value, ok := solarSystem["Pluto"]
	fmt.Println(value, ok)

	if value, ok := solarSystem["Mercury"]; ok {
		fmt.Println(value)
	}

	for key, value := range solarSystem {
		fmt.Println(key, value)
	}

	for _, value := range solarSystem {
		fmt.Println(value)
	}
}
