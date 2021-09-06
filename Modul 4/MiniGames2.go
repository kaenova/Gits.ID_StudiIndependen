package main

import "fmt"

func main() {
	var (
		jumlah int
		angka  int = 1
	)

	fmt.Print("Total Deret : ")
	fmt.Scanf("%d", &jumlah)

	if jumlah < 0 {
		fmt.Println("Total Deret tidak boleh negatif")
	} else {
		for i := 0; i < jumlah; i++ {
			fmt.Printf("%d ", angka)
			if i%2 == 0 {
				angka = angka + 4
			} else {
				angka = angka - 2
			}
		}
	}
	fmt.Printf("\n")
}
