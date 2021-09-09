package main

import "fmt"

func main() {
	var bruh = cal(1, 2, 3, 4, 5, 6)
	fmt.Println(bruh)
}

func cal(numbers ...int) (total int) {
	total = 0
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)
	for _, num := range numbers {
		total += num
	}
	return
}
