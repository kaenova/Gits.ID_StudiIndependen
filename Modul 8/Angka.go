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
	fmt.Println("Jumlah angka :", JumlahAngka(kata))
}

func JumlahAngka(kata string) uint {
	var hasil uint
	current := byte(kata[0])
	if (current >= byte('0')) && (current <= byte('9')) {
		hasil = 1
	} else {
		hasil = 0
	}
	if len(kata) == 1 {
		return hasil
	} else {
		kata_baru := kata[1:]
		return hasil + JumlahAngka(kata_baru)
	}
}
