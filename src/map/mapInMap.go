package main

import "fmt"

func main() {
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022e+23,
			"orbitalPeriod": 87.969,
		},

		"Venus": map[string]float32{
			"meanRadius":    612.7,
			"mass":          2.3022e+23,
			"orbitalPeriod": 15.969,
		},

		"Mars": map[string]float32{
			"meanRadius":    3398.7,
			"mass":          1.3022e+23,
			"orbitalPeriod": 994.969,
		},
	}

	fmt.Println(terrestrialPlanet)

	fmt.Println(terrestrialPlanet["Mars"]["mass"])
}
