/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import "fmt"

func main() {
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)
	fmt.Println("Jumlah huruf non kapital :", JumlahAngka(kata))
}

func JumlahAngka(kata string) uint {
	var jumlah uint = 0
	for i := 0; i < len(kata); i++ {
		if (kata[i] >= 48) && (kata[i] <= 57) {
			jumlah++
		}
	}
	return jumlah
}
