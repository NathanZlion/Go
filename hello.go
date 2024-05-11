package main

import "fmt"

func main() {
	myBill := newBill("Lunch");
	fmt.Println(myBill);
	// fmt.Println(myBill);
	// updateName(&myBill, "Dinner");
	// fmt.Println(myBill);

	// addItem(&myBill, "Genfo", 5.3);
	// fmt.Println(myBill);
	myBill.addItem("Fried Chicken & Rice", 430.0);
	myBill.addItem("Coke", 20.0);

	fmt.Println(myBill.format());
}
