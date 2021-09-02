package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var (
		harga float32
		quantity int
	)

	in := bufio.NewReader(os.Stdin)

	fmt.Print("Nama Barang : ")
	namaBarang, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	namaBarang = strings.ReplaceAll(namaBarang, "\n", "")

	fmt.Print("Harga Barang : ")
	fmt.Scanf("%f", &harga)

	fmt.Print("Quantity Barang : ")
	fmt.Scanf("%d", &quantity)

	total := harga * float32(quantity)

	fmt.Printf("Barang %s berjumlah %d memiliki harga satuan %.3f dengan harga total %.3f", namaBarang, quantity, harga, total)

}