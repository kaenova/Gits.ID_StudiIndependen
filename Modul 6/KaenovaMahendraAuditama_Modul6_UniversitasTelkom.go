package main

import (
	"fmt"
	"strconv"
)

var (
	data    [][4]string // ID, NAMA, SCORE
	menu    bool        = true
	pilihan int         = -1
)

func main() {

	for menu {
		fmt.Println("Menu")
		fmt.Println("1. Masukkan data game")
		fmt.Println("2. Hapus data berdasarkan ID")
		fmt.Println("3. Tampilkan seluruh data dan jumlah")
		fmt.Println("4. Cari berdasarkan judul")
		fmt.Println("5. Tampilkan 3 game favorit")
		fmt.Println("6. Tampilkan rating diatas 4.0")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&pilihan)
		fmt.Println("================")

		switch pilihan {
		case 1:
			var (
				input_nama       string
				input_rating     float32
				input_id         string
				input_keterangan string
				ulang            bool = true
			)
			for ulang {
				for ulang {
					fmt.Print("Masukkan ID: ")
					fmt.Scanln(&input_id)
					if len(input_id) > 7 {
						fmt.Println("Input ID tidak sesuai format, harap masukkan dengan panjang maksimal 7 huruf")
					} else {
						ulang = false
					}
				}

				// Cek id sudah pernah dipakai atau belum
				for i := 0; i < len(data); i++ {
					if input_id == data[i][0] {
						ulang = true
						fmt.Println("ID sudah pernah dibuat pada game", data[i][1])
						break
					}
				}
			}

			fmt.Print("Masukkan nama game: ")
			fmt.Scanln(&input_nama)

			ulang = true
			for ulang {
				fmt.Print("Masukkan rating: ")
				fmt.Scanln(&input_rating)
				if (input_rating <= 5.0) && (input_rating >= 0.0) {
					ulang = false

				} else {
					fmt.Println("Rating tidak valid (0.0 - 5.0)")
				}
			}
			input_keterangan = "not_defined"
			if input_rating > 4.0 {
				input_keterangan = "good"
			} else if input_rating >= 2.0 {
				input_keterangan = "average"
			} else if input_rating >= 0 {
				input_keterangan = "poor"
			}
			// Masukkan ke data utama
			var data_temp = [4]string{input_id, input_nama, fmt.Sprintf("%.1f", input_rating), input_keterangan}
			data = append(data, data_temp)
			fmt.Println("Data berhasil dimasukkan")

		case 2:
			var (
				input_id    string
				index       int = -1
				kiri, kanan [][4]string
			)
			if len(data) > 0 {
				fmt.Print("Masukkan ID data yang ingin dihapus : ")
				fmt.Scanln(&input_id)
				// Cari index by id
				for i := 0; i < len(data); i++ {
					if input_id == data[i][0] {
						index = i
					}
				}

				if index != -1 {
					// Ditemukan ID
					kiri = data[:index]
					kanan = data[index+1:]
					data = append(kiri, kanan...)
				} else {
					// ID Tidak ditemukan
					fmt.Println("ID Tidak ditemukan")
				}

			} else {
				fmt.Println("Data kosong")
			}
		case 3:
			if len(data) > 0 {
				fmt.Printf("|%7s|%15s|%8s|%12s|\n", "ID", "Nama", "Score", "Keterangan")
				fmt.Printf("|%7s|%15s|%8s|%12s|\n", "-------", "---------------", "--------", "------------")
				for i := 0; i < len(data); i++ {
					fmt.Printf("|%7s|%15s|%8s|%12s|\n", data[i][0], data[i][1], data[i][2], data[i][3])
				}
				fmt.Println("Jumlah data :", len(data))
			} else {
				fmt.Println("Data kosong")
			}
		case 4:
			var (
				ketemu     bool = false
				input_nama string
			)
			fmt.Print("Masukkan nama game : ")
			fmt.Scanln(&input_nama)
			for i := 0; i < len(data); i++ {
				if data[i][1] == input_nama {
					ketemu = true
					fmt.Println("Data ditemukan")
					fmt.Printf("|%7s|%15s|%8s|%12s|\n", "ID", "Nama", "Score", "Keterangan")
					fmt.Printf("|%7s|%15s|%8s|%12s|\n", "-------", "---------------", "--------", "------------")
					fmt.Printf("|%7s|%15s|%8s|%12s|\n", data[i][0], data[i][1], data[i][2], data[i][3])
					break
				}
			}
			if !ketemu {
				fmt.Println("Data tidak ditemukan pada game", input_nama)
			}
		case 5:
			var (
				idx_nilai_terbesar int
			)

			if len(data) > 0 {
				fmt.Println("Menampilkan 3 game favorit berdasarkan score")

				// Buat data copy agar aman
				data_copy := make([][4]string, len(data))
				copy(data_copy, data)

				// sort data_copy by rating using ParseFloat
				for i := 0; i < len(data_copy)-1; i++ {
					idx_nilai_terbesar = i
					for j := i + 1; j < len(data_copy); j++ {
						num1, err := strconv.ParseFloat(data_copy[idx_nilai_terbesar][2], 2)
						if err != nil {
							fmt.Println("Error when converting to float of", data_copy[idx_nilai_terbesar][2])
						}
						num2, err := strconv.ParseFloat(data_copy[j][2], 2)
						if err != nil {
							fmt.Println("Error when converting to float of", data_copy[j][2])
						}
						if num1 < num2 {
							idx_nilai_terbesar = j
						}
					}
					// penukaran
					temp := data_copy[i]
					data_copy[i] = data_copy[idx_nilai_terbesar]
					data_copy[idx_nilai_terbesar] = temp
				}

				fmt.Printf("|%7s|%15s|%8s|%12s|\n", "ID", "Nama", "Score", "Keterangan")
				fmt.Printf("|%7s|%15s|%8s|%12s|\n", "-------", "---------------", "--------", "------------")
				for i := 0; i < 3; i++ {
					fmt.Printf("|%7s|%15s|%8s|%12s|\n", data_copy[i][0], data_copy[i][1], data_copy[i][2], data_copy[i][3])
				}

			} else {
				// Kalau misalkan datanya kosong
				fmt.Println("Data kosong")
			}
		case 6:
			var counter int = 0
			if len(data) > 0 {
				fmt.Println("Menampilkan game yang memiliki score >4.0")
				fmt.Printf("|%7s|%15s|%8s|%12s|\n", "ID", "Nama", "Score", "Keterangan")
				fmt.Printf("|%7s|%15s|%8s|%12s|\n", "-------", "---------------", "--------", "------------")
				for i := 0; i < len(data); i++ {
					nilai, err := strconv.ParseFloat(data[i][2], 2)
					if err != nil {
						fmt.Println("Error when converting to float of", data[i][2])
					}
					if nilai > 4.0 {
						fmt.Printf("|%7s|%15s|%8s|%12s|\n", data[i][0], data[i][1], data[i][2], data[i][3])
						counter++
					}
				}
				fmt.Println("Jumlah data yang memiliki score >4.0:", counter)
			} else {
				fmt.Println("Data kosong")
			}
		case 7:
			menu = false
		default:
			fmt.Println("Pilihan tidak dikenali, kembali ke menu")
		}
		fmt.Println("===================")
		pilihan = -1

	}

}
