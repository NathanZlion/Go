package main

import "fmt"

func sayHi(name string) {
	fmt.Printf("Hi there %v.\n", name)
}


// Passing functions as function argument
func cycleNames (names []string, function func(string)) {
	for _, name	:= range names {
		function(name);
	}
}

func main() {
	fmt.Println("Hello World")
	names := []string {"Nathnael", "Alebachew", "Yonas"}
	cycleNames(names, sayHi);
}

