package main

import "fmt"

func main() {

	a := []string{"Kaenova", "Mahendra", "Auditama"}
	fmt.Println("a :", a)
	b := []string{"Auditama", "Kaenova", "Mahendra"}
	a = append(a, "Halo")
	fmt.Println("b :", b)
	fmt.Println("a :", a)
}
