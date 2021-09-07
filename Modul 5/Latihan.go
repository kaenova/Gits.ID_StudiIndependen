package main

import (
	"fmt"
)

func main() {
	var (
		numA, numB, numMenu int
		ulang_input         string
		ulang               bool = true
	)

	for ulang {
		fmt.Println("========================")
		fmt.Print("Masukan Angka ke 1 : ")
		fmt.Scan(&numA)
		fmt.Print("Masukan Angka ke 2 : ")
		fmt.Scan(&numB)
		fmt.Println("========================")
		fmt.Println("Pilih menu dibawah :")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("========================")
		fmt.Print("Pilih menu angka diatas : ")
		fmt.Scan(&numMenu)

		if (numMenu < 0) || (numMenu > 4) {
			fmt.Println("Maaf pilihan tidak dikenali")
		} else {
			switch numMenu {
			case 1:
				fmt.Println("Hasil penjumlahan", numA, "+", numB,
					"=", numA+numB)
			case 2:
				fmt.Println("Hasil pengurangan", numA, "-", numB,
					"=", numA-numB)
			case 3:
				fmt.Println("Hasil pengurangan", numA, "*", numB,
					"=", numA*numB)
			case 4:
				fmt.Println("Hasil pengurangan", numA, "/", numB,
					"=", float32(numA)/float32(numB))
			}
		}

		// Minta input ulang
		fmt.Print("Ingin diulangi lagi (y/N)? ")
		fmt.Scanf("\n%s", &ulang_input)
		if ulang_input == "N" {
			ulang = false
		}
	}
}
