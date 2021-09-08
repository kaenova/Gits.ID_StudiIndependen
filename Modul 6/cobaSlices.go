package main

import "fmt"

func main() {
	var (
		kiri, kanan []int
		index       uint = 4
	)

	var a = []int{1, 2, 3, 4, 5}
	kiri = a[:index]
	kanan = a[(index + 1):]
	a = append(kiri, kanan...)
	fmt.Println(a)
}
