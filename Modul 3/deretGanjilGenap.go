/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import "fmt"

func main(){

	// Jumlah Inputan User
	var (
		jumlah, i, counter uint
	)
	fmt.Print("Masukkan jumlah unsur deret : ")
	fmt.Scanf("%d", &jumlah)

	// Menampilkan deret ganjil
	fmt.Println("Deret Ganjil :")
	counter = 0
	i = 0
	for i<jumlah {
		if counter % 2 != 0 {
			fmt.Printf(" %d", counter)
			i++
		}
		counter++
	}
	fmt.Printf("\n")

	// Menampilkan deret genap
	fmt.Println("Deret Genap :")
	counter = 0
	i = 0
	for i<jumlah {
		if counter % 2 == 0 {
			fmt.Printf(" %d", counter)
			i++
		}
		counter++
	}
	fmt.Printf("\n")
}