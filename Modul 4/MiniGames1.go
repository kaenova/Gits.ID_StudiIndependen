package main

import (
	"fmt"
)

func main() {

	var (
		jumlah  int
		total   int = 0
		counter int = 2
	)

	fmt.Print("Total Deret : ")
	fmt.Scanf("%d", &jumlah)

	if jumlah < 0 {
		fmt.Println("Total Deret tidak boleh negatif")
	} else {
		for i := 0; i < jumlah; i++ {
			total = total + counter
			if i+1 != jumlah {
				fmt.Printf("%d + ", counter)
			} else {
				fmt.Printf("%d =", counter)
			}
			counter = counter + 2
		}
		fmt.Printf(" %d\n", total)
	}
}
