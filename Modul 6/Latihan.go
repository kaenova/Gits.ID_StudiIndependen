package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		Menu
		1. Input Data
		2. Menghapus Data by Index
		3. Tampilkan Semua Data
		4. Tampilkan dikon >15%
		5. Menampilkan harga setelah diskon paling murah
		6. Cari by Nama
		7. Keluar
		Masukkan nomor menu :
	*/

	// Inisialisasi Variable
	var (
		ulang        bool = true
		pilihan      int
		kode         []uint
		nama         []string
		harga        []float32
		diskon       []float32 // dalam bentuk decimal
		harga_diskon []float32
	)
	in := bufio.NewReader(os.Stdin)

	first := true

	for ulang {
		fmt.Println("Menu")
		fmt.Println("1. Input Data")
		fmt.Println("2. Menghapus Data by Index")
		fmt.Println("3. Tampilkan Semua Data")
		fmt.Println("4. Tampilkan dikon >15%")
		fmt.Println("5. Menampilkan harga setelah diskon paling murah")
		fmt.Println("6. Cari by Nama")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan nomor menu : ")
		if first {
			fmt.Scanf("%d", &pilihan)
			first = false
		} else {
			fmt.Scanf("\n%d", &pilihan)
		}
		fmt.Println("========================")

		switch pilihan {
		case 1:
			/*
				Input Data
				Melakukan input data dan update harga terkecil
			*/
			fmt.Println("Inputting Data")
			var (
				nama_dicari   string
				kode_barang   uint
				harga_barang  float64
				diskon_barang float64
				ulang_input   bool = true
			)
			for ulang_input {
				fmt.Print("Kode barang : ")
				fmt.Scanf("\n%d", &kode_barang)
				fmt.Println(kode, kode_barang)
				if len(kode) > 0 {
					for i := 0; i < len(kode); i++ {
						if kode_barang == kode[i] {
							fmt.Println("Kode sudah terpakai pada barang index ke", i)
							ulang_input = true
						} else {
							ulang_input = false
							kode = append(kode, kode_barang)
						}
					}
				} else {
					ulang_input = false
					kode = append(kode, kode_barang)
				}

				if !ulang_input {
					// Nama Barang
					fmt.Print("Nama barang : ")
					fmt.Scanf("\n%s", &nama_dicari)
					nama_dicari = strings.ReplaceAll(nama_dicari, "\n", "")
					nama_dicari = strings.ToLower(nama_dicari)
					nama = append(nama, nama_dicari)
					ulang_input = false
				}

				if !ulang_input {
					// Harga Barang
					fmt.Print("Harga barang : ")
					fmt.Scanf("\n%f", &harga_barang)
					if harga_barang < 0 {
						ulang_input = true
					} else {
						harga_barang, err := strconv.ParseFloat(fmt.Sprintf("%.2f", harga_barang), 2)
						if err != nil {
							fmt.Println("Sistem error, perubahan harga barang")
						}
						harga = append(harga, float32(harga_barang))
					}
				}

				if !ulang_input {
					// Diskon
					ulang_input = true
					for ulang_input {
						fmt.Print("Masukkan diskon (desimal) : ")
						fmt.Scanf("\n%f", &diskon_barang)
						if (diskon_barang <= 1) && (diskon_barang >= 0) {
							ulang_input = false
							diskon_barang, err := strconv.ParseFloat(fmt.Sprintf("%.2f", diskon_barang), 2)
							if err != nil {
								fmt.Println("Sistem error, perubahan harga barang")
							}
							diskon = append(diskon, float32(diskon_barang))
							harga_diskon = append(harga_diskon, float32(diskon_barang*harga_barang))
						} else {
							fmt.Println("Diskon tidak valid, masukkan dalam bentuk desimal")
						}
					}
				}
				if ulang_input {
					fmt.Println("=================")
					fmt.Println("Mengulangi Input")
				} else {
					fmt.Println(" Data Berhasil Dimasukkan ")
				}
			}
		case 2:
			/*
				Menghapus Data by Index
				Update harga terkecil
			*/

		case 3:
			var patokan_angka int = len(kode)

			if patokan_angka != len(nama) || patokan_angka != len(harga) || patokan_angka != len(diskon) || patokan_angka != len(harga_diskon) {
				fmt.Println("Data corrupted, exiting program")
				ulang = false
				break
			}
			for i := 0; i < len(kode); i++ {
				fmt.Println("[", i, "]", "Barang ke", i+1)
				fmt.Println("Kode barang :", kode[i])
				fmt.Println("Nama barang :", nama[i])
				fmt.Println("Harga barang :", harga[i])
				fmt.Println("Diskon barang :", diskon[i])
				fmt.Println("Harga setelah diskon :", harga_diskon[i])
			}
		case 4:
			var patokan_angka int = len(kode)
			fmt.Println(kode, nama, harga, diskon, harga_diskon)
			if (patokan_angka != len(nama)) || (patokan_angka != len(harga)) || (patokan_angka != len(diskon)) || (patokan_angka != len(harga_diskon)) {
				fmt.Println("Data corrupted, exiting program")
				ulang = false
				break
			}
			var j int = 1
			for i := 0; i < len(kode); i++ {
				if diskon[i] > 0.15 {
					fmt.Println("Barang ke", j)
					fmt.Println("Kode barang :", kode[i])
					fmt.Println("Nama barang :", nama[i])
					fmt.Println("Harga barang :", harga[i])
					fmt.Println("Diskon barang :", diskon[i])
					fmt.Println("Harga setelah diskon :", harga_diskon[i])
					j++
				}
			}
		case 5:

		case 6:
			fmt.Print("Nama barang : ")
			nama_dicari, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			nama_dicari = strings.ReplaceAll(nama_dicari, "\n", "")
			nama_dicari = strings.ToLower(nama_dicari)
			fmt.Println(nama_dicari)
			for i := 0; i < len(kode); i++ {
				if nama_dicari == strings.ToLower(nama[i]) {
					fmt.Println("Kode barang :", kode[i])
					fmt.Println("Nama barang :", nama[i])
					fmt.Println("Harga barang :", harga[i])
					fmt.Println("Diskon barang :", diskon[i])
					fmt.Println("Harga setelah diskon :", harga_diskon[i])
				}
			}
		case 7:
			fmt.Println("Exiting program")
			ulang = false
		default:
			fmt.Println("Pilihan tidak dimengerti, kembali ke menu")
		}
		fmt.Println("===================")
	}
}
