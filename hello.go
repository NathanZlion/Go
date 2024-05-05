package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	fmt.Println("Hello World")

	// ----------------- strings ----------------- //
	// strings.Contains
	greeting := "Hello there, how are you?";
	fmt.Println(strings.Contains(greeting, "Hello"));

	// strings.ReplaceAll
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"));

	/* The functions do not alter the original stirng. */

	// strings.ToUpper
	fmt.Println(strings.ToUpper(greeting));

	// strings.Index : first occurrence
	fmt.Println(strings.Index(greeting, "llo"));

	// string.Split
	// string.Join
	fmt.Println(strings.Join(strings.Split(greeting, " "), "|"));

	// -----------------sort----------------- //
	// sort.Ints
	ages := []int {2, 4, 3, 1};
	sort.Ints(ages);
	fmt.Println(ages);

	/* You must notice here that this method has modified the original array. */

	// sort.SearchInts : returns the index of the first occurrence in the array.
	// If the elemeent does not exist it returns the length of the slice.
	fmt.Println(sort.SearchInts(ages, 3));

	names := []string {"nathnael", "yoseph", "Abel"};
	sort.Strings(names)
	fmt.Println(names);

	fmt.Println(sort.SearchStrings(names, "nathnael"));
}

