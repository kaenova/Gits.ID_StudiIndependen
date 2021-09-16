package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var (
		input, output string
		err           error
	)
	fmt.Print("Masukkan input : ")
	fmt.Scanln(&input)
	output, err = Hasil(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("Output :", output)
}

func Hasil(kal string) (string, error) {
	// String lower agar if nya cuman di bagian huruf kecil
	kal = strings.ToLower(kal)
	output := ""
	for i := len(kal) - 1; i >= 0; i-- {
		if kal[i] >= byte('a') && kal[i] <= byte('z') {
			output = output + string(kal[i])
		}
	}
	// Jika tidak terdapat huruf maka kembalikan error
	if len(output) == 0 {
		return "", errors.New("Harus terdapat string")
	}
	// Ambil setiap kata menjadi huruf besar
	output = strings.Title(output)
	return output, nil
}
