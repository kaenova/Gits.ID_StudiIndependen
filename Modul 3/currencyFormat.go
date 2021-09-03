/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){

	var (
		nilai, mod int
		final string
		minus, first_conversion bool
	)
	// Currency Initialization
	minus = false
	final = ",00"
	first_conversion = true

	//Input
	fmt.Print("Masukkan jumlah nilai : ")
	fmt.Scanf("%d", &nilai)

	// Minus check
	if nilai < 0 {
		minus = true
		nilai = nilai * -1
	}

	// Konversi nilai
	for nilai > 0 {
		mod = nilai % 1000
		nilai = nilai / 1000

		if first_conversion {
			first_conversion = false
		} else {
			final = "."+final
		}

		if mod == 0 {
			final = "000" + final
		} else if mod < 10 {
			// Kurang dari sepuluh
			final = strconv.Itoa(mod)+final
			if nilai > 10 {
				// Kalau misalkan ada yang masih bisa dibagi lagi
				final = "00" + final
			}
		} else if mod / 10 < 10 {
			// Kurang dari 100
			final = "0"+strconv.Itoa(mod)+final
		} else {
			final = strconv.Itoa(mod)+final
		}

	}

	// Finalisasi
	final = "Rp " + final
	if minus {
		final = "- " + final
	}

	// Output
	fmt.Println("Hasil konversi :", final)
}