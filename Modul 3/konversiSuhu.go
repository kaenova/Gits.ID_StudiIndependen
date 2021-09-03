/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import "fmt"

func main(){
	var (
		celcius float32
	)
	fmt.Print("Masukkan suhu dalam celcius : ")
	fmt.Scanf("%f", &celcius)

	// Konversi ke fahrenheit
	fahrenheit := (celcius * (9.0/5.0)) + 32
	fmt.Println("Konversi suhu ke fahrenheit :", fahrenheit)

	// Koversi ke kelvin
	kelvin := (celcius + 273.15)
	fmt.Println("Konversi suhu ke kelvin :", kelvin)
}