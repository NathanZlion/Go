package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateMap(x map[string]int) {
	x["nathnael"] = 3;
}

func updateNameByPointer (name *string) {
	*name = "tewodros";
}

func main() {
	/*
		When passed into functions go makes a copy of the variables.
		So pass by value.
	*/

	/* Group 1 - Non-Pointer Values - int, strings, bools, floats, arrays, structs. */
	name := "tifa"
	updateName(name)
	fmt.Println(name)
	// Output: tifa. Because the function updateName makes a copy of the variable name.


	/* Group 2 - Pointer Wrapper Values - slices, maps, functions. */

	map_ := map[string]int {
		"nathnael": 1,
		"tewodros": 2,
	}

	updateMap(map_);
	fmt.Println(map_)
	// Output: map[nathnael:3 tewodros:2], updates


	/*
	We could want to pass a Non-Pointer Value by reference, so that the function when it copies is
	actually copying the reference to the value.
	*/
	fmt.Printf("The memory address of name is %v \n", &name);
	// Output: The memory address of name is 0xc0000101e0. And obviously he memory address changes every time the program is run.
	fmt.Printf("and the type of the memory address  is %T \n", &name);
	// Output:  and the type of the memory address  is *string

	m := &name;
	fmt.Println(m == &name); // true. Both hold the pointer to the same memory address.

	// We can dereference the pointer to get the value. Just place a * (asterisk) before the variable name.
	fmt.Println(*m);

	// Now let's see how we could update non pointer variable types.
	my_name := "nathnael";
	fmt.Printf("My name was %s ", my_name);
	updateNameByPointer(&my_name);
	fmt.Printf("but now it is %s.", my_name);

	// So remember:
	// & gives the memory address
	// * gives the value stored at the memory address for a pointer.

}
