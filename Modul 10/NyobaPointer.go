package main

import "fmt"

func main() {
	var a = 5
	var p = &a                              // p holds variable a's memory address
	fmt.Printf("Address of var a: %p\n", p) // isi dari p itu, address memori
	fmt.Printf("Value of var a: %v\n", *p)  // ngambil isi yang ada dari address p

	// Let's change a value (using the initial variable or the pointer)
	*p = 3 // using pointer

	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)
	a = 10 // using initial var
	fmt.Printf("Value of var a: %v\n", *p)

}
