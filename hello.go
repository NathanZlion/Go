package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	map_ := map[string]float64 {
		"nathnael": 2.3,
		"dawit": 1.3,
		"luigi": 21.4,
	}

	fmt.Println(map_);
	fmt.Println(map_["nathnael"]);

	// traversing the map
	for key, value := range map_ {
		fmt.Printf("key : %s, Value : %f\n", key, value);
	}
}

