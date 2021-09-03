/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import "fmt"

func main(){
	var panjang, lebar float64

	// Input
	fmt.Print("Masukkan panjang : ")
	fmt.Scanf("%f", &panjang)

	fmt.Print("Masukkan lebar : ")
	fmt.Scanf("%f", &lebar)

	luas := panjang * lebar

	keliling := (2*panjang) + (2*lebar)

	fmt.Println("Luas bidang :", luas)
	fmt.Println("Keliling bidang :", keliling)
}