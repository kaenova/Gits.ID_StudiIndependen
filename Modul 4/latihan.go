package main

import (
	"fmt"
)

func main() {
	var numA, numB, numMenu int
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
	if numMenu == 1 {
		fmt.Println("Hasil penjumlahan", numA, "+", numB,
			"=", numA+numB)
	} else if numMenu == 2 {
		fmt.Println("Hasil pengurangan", numA, "+", numB,
			"=", numA-numB)
	} else if numMenu == 3 {
		fmt.Println("Hasil pengurangan", numA, "*", numB,
			"=", numA*numB)
	} else if numMenu == 4 {
		fmt.Println("Hasil pengurangan", numA, "/", numB,
			"=", float32(numA)/float32(numB))
	} else {
		fmt.Println("Pilihan operator tidak dikenali")
	}
}
