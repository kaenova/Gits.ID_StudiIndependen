package main

import "fmt"

func main(){
	a := 321521
	count := 0
	for a > 0 {
		fmt.Println("Iterasi ", count+1)
		fmt.Println(a / 1000)
		fmt.Println(a % 1000)
		a = a /1000
		count++
	}
}