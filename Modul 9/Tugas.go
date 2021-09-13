package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Inisialisasi Menu utama
var menu = []string{"Tahu", "Tempe", "Sambal", "Nasi", "Lele", "Ayam"}

func main() {
	// Sorting for better validation
	sort.Slice(menu, func(i, j int) bool { return menu[i] < menu[j] })

	// Print Menu Utama
	fmt.Println("Toko makanan Indonesia")
	fmt.Println("=======================")
	fmt.Printf("|%-3s|%8s|\n", "No.", "Menu")
	fmt.Printf("%14s\n", "=============")
	for i, makanan := range menu {
		fmt.Printf("|%-3d|%8s|\n", i+1, makanan)
	}

	// Input Pesanan
	var (
		pesanan             []string
		lanjut              bool = true
		temp_pesanan        string
		temp_pilihan_lanjut string
	)
	for lanjut {
		fmt.Print("Masukkan manu pesanan anda dalam huruf ( eg: Tahu ) : ")
		fmt.Scanln(&temp_pesanan)
		if valid, err := Validator(temp_pesanan, menu); err != nil {
			// Print jika ada error
			panic(err)
		} else {
			if !valid {
				fmt.Println("Pesanan anda tidak tersedia")
			} else {
				pesanan = append(pesanan, temp_pesanan)
			}
		}
		fmt.Print("Lanjutkan memesan ? (Y/T) : ")
		fmt.Scanln(&temp_pilihan_lanjut)
		if strings.ToLower(strings.TrimSpace(temp_pilihan_lanjut)) == "t" {
			lanjut = false
		}
	}

	// Output pesanan
	fmt.Printf("\n%14s\n", "==============")
	fmt.Printf("|%11s|\n", "Pesanan Anda")
	fmt.Printf("|%-3s|%8s|\n", "No.", "Menu")
	fmt.Printf("%14s\n", "==============")
	for i, makanan := range pesanan {
		fmt.Printf("|%-3d|%8s|\n", i+1, makanan)
	}

}

func Validator(input_makanan string, menu_utama_sorted []string) (bool, error) {
	/*
		Editor note -Kaenova :
		Masih tidak yakin ini paling efektif. Awalnya inign dilakukan binary search
		secara langsung, tetapi kepikiran harus dilakukan pengecekkan apakah sudah
		terurut atau belum, masalahnya pengecekkan itu sptnya dilakukan secara linear
		time, O(n). Jadi sama aja.

		Disini aku juga mencoba membuat error, cuman masih ragu kapan digunakan
		error sama kapan bilang ga valid. Menurut Aku, definisi valid disini itu
		terhadap input_makanan(-nya). Artinya valid jika input_makanan itu
		tersedia di menu_utama_sorted, dan tidak valid jika input_makanan
		tidak ada di menu_utama_sorted. Sisanya akan saya buat error. Apakah benar
		pembuatan error dengan logika seperti itu?

		Karena aku bisa aja yang error saya buat tidak valid. Mohon bantuannya kak.
	*/

	// Check if menu is not empty
	if len(menu_utama_sorted) <= 0 {
		return false, errors.New("Menu kosong")
	}

	// Check if input is empty
	if len(strings.TrimSpace(input_makanan)) == 0 {
		return false, nil
	}

	// Check Sorted slice
	if !(sort.SliceIsSorted(menu_utama_sorted, func(i, j int) bool {
		return menu_utama_sorted[i] < menu_utama_sorted[j]
	})) {
		return false, errors.New("Menu belum dilakukan sorting")
	}

	// Binary search
	var (
		kiri   int = 0
		kanan  int = len(menu) - 1
		tengah     = (kiri + kanan) / 2
	)

	input_makanan = strings.ToLower(input_makanan)

	for kiri < kanan {
		if (strings.ToLower(menu[tengah]) == input_makanan) ||
			(strings.ToLower(menu[kiri]) == input_makanan) ||
			(strings.ToLower(menu[kanan]) == input_makanan) {
			return true, nil
		} else if strings.ToLower(menu[tengah]) < input_makanan {
			kiri = tengah + 1
		} else if strings.ToLower(menu[tengah]) > input_makanan {
			kanan = tengah - 1
		}
		tengah = (kiri + kanan) / 2
	}

	// Not found on binary search
	return false, nil
}
