package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// In go for plays the dual role, both as a
	// 'for' and as a 'while'

	x := 0;
	for x < 5 {
		fmt.Printf("The value of x is %d\n", x);
		x++;
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("The value of i is %d\n", i);
	}

	// Looping through a slice
	numbers := []int {1, 3, 5, 7};

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %v, value %v\n", i, numbers[i]);
	}

	// range, the python equivalent to enumerate
	for index, value := range(numbers) {
		fmt.Printf("Index %v, value %v\n", index, value);
	}

	/*
	If you don't want to use a variable let's say the index in this case, you can 
	use _ (underscore) as the variable name. This is not a suggestion it's an order
	because go doesn't compile with unused variables in it.
	 */
}

