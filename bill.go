package main

import "fmt"

// ! Beware there are no commas between the fields
type bill struct {
	name string
	items map[string]float64
	tip float64
}

// creator function
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b;
}

// Receiver functions.
/* The receiver function is a function that is attached to a type.
   We can use it like obj.method() like in other languages. */

func (b *bill) addItem(name string, value float64) {
	b.items[name] = value;
}

func (b bill) format() string {
	var fs string;
	fs += fmt.Sprintf ("%v\n", "----------------------");
	fs += "Bill breakdown: \n";
	var total float64 = 0;

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v  %-5v ETB\n", k + ":", v);
		total += v;
	}
	fs += fmt.Sprintf ("%v\n", "----------------------");

	fs += fmt.Sprintf("%-25v  %-5v ETB\n", "tip:", b.tip);
	fs += fmt.Sprintf("%-25v  %-5v ETB\n", "total:", total + b.tip);

	return fs;
}
