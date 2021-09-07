/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		jenis string
		baris int
	)
	fmt.Print("Masukkan jenis segitiga ? (Terbalik / Normal) : ")
	fmt.Scanf("%s", &jenis)
	fmt.Print("Masukkan Jumlah Baris ? : ")
	fmt.Scanf("\n%d", &baris)

	jenis = strings.ToLower(jenis)

	jumlah_max_bintang := (2 * baris) - 1
	tengah := jumlah_max_bintang / 2
	jumlah_max_spasi := tengah

	switch jenis {
	case "normal":
		for i := 0; i < baris; i++ {
			for j := 0; j < jumlah_max_spasi; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < (2*(i+1) - 1); j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")
			jumlah_max_spasi--
		}
	case "terbalik":
		for i := baris; i > 0; i-- {
			for j := 0; j < tengah-jumlah_max_spasi; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < (2*(i) - 1); j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")
			jumlah_max_spasi--
		}
	}
}
