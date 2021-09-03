/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	// Input Kalimat
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat: ")
	kalimat, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	kalimat = strings.ReplaceAll(kalimat, "\n", "")

	// Menghitung jumlah karakter
	var mapping = make(map[byte]int)
	for i := 0 ; i<len(kalimat); i++{
		mapping[kalimat[i]] ++
	}

	// Output
	fmt.Println("Jumlah karakter yang diinput :", len(kalimat))
	fmt.Println("Karakter yang Sama dengan Jumlahnya")
	for key, value := range mapping {
		if value > 1 {
			fmt.Println("'",string(key),"'", ":", uint32(value))
		}
	}
}