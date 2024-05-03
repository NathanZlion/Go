package main

import "fmt"

func main() {

	// Print : does not add a newline
	fmt.Print("Hello ")
	fmt.Print("World\n")

	// Println : adds a newline \n
	fmt.Println("Hello World")

	age := 22;
	name := "Nathnnael";

	fmt.Println("My name is", name, "and I am ", age, "years old.");

	// Printf : Formatted Strings
	fmt.Printf("My name is %v and I am %v years old.\n", name, age);

	// %q : adds quote to strings.
	fmt.Printf("My name is %q and I am %q years old.\n", name, age); 

	// %T : type
	fmt.Printf("My name is of type %T and age is of type %T. \n", name, age);

	// %f : floating number
	fmt.Printf("The score is %f\n", 33.2);
	// Output : The score is 33.200000

	// %0.#f : floating number but you can specify the decimal point precision.
	fmt.Printf("The score is %0.1f\n", 33.2);
	// Output : The score is 33.2

	// Sprintf : save formatted string
	var output = fmt.Sprintf("My name is %v and I am %v years old.\n", name, age);
	fmt.Printf("Output => %s", output);

	// Check this link out to get a list of all the available such characters
	// https://pkg.go.dev/fmt
}

