package main

import (
	"fmt"
)

func main() {
	a := "azAZ09"
	fmt.Println(a[0], a[1], a[2], a[3], a[4], a[5])
	fmt.Println(byte('a'))
	b := "abc"
	fmt.Printf("%s", b[1:])
}
