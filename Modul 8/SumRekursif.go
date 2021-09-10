/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import "fmt"

func main() {
	var awal float32
	var rek_max int
	fmt.Print("Masukkan angka awal : ")
	fmt.Scanln(&awal)
	fmt.Print("Masukkan banyaknya rekursif : ")
	fmt.Scanln(&rek_max)
	fmt.Print("Sum = ")
	Rekursif(awal, 0, 0, rek_max)
}

func Rekursif(awal, jumlah float32, current, rek_max int) {
	if current < rek_max {
		temp := awal / float32(current+1)
		if current == rek_max-1 {
			fmt.Printf(" %.2f =", temp)
		} else {
			fmt.Printf(" %.2f +", temp)
		}
		jumlah = jumlah + temp
		Rekursif(awal, jumlah, current+1, rek_max)
	} else {
		fmt.Printf(" %.2f", jumlah)
	}
}
