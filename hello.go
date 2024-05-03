package main

import "fmt"

func main() {
	// Arrays in Go and also other languages as you know it have fixed size.
	// Explicit initialization
	var age [3]int  = [3]int {12, 3, 3};
	var also_age = [3]int {12, 3, 3}; // simpler way
	fmt.Println(age);
	fmt.Println(also_age);

	// len : Length of an array
	fmt.Println(age, "Length =>" , len(age));

	// As you have seen arrays are of fixed size. Once initialized you cannot change the size.
	// In python you know list right. Thq equivalent of list in Go is slice.
	// Slices are dynamic arrays. They can grow and shrink. Under the hood slices are implemented as arrays.
	var scores = []int {1, 2, 3};
	fmt.Println(scores);

	// Append to a slice
	scores = append(scores, 4);
	fmt.Println(scores, "Length =>" , len(scores));
	fmt.Printf("The type of scores is %T\n", scores);

	// Slice range
	rangeOne := scores[0:2];
	fmt.Println(rangeOne, "Length =>" , len(rangeOne));
	fmt.Printf("The type of rangeOne is %T\n", rangeOne);
}

