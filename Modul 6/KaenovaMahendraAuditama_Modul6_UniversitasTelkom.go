package main

import (
	"fmt"
	"strconv"
)

var (
	data    [][3]string // ID, NAMA, SCORE
	menu    bool        = true
	pilihan int
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
				input_nama   string
				input_rating float32
				input_id     string
				ulang        bool = true
			)
			for ulang {
				fmt.Print("Masukkan ID: ")
				fmt.Scanln(&input_id)
				ulang = false
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

			// Masukkan ke data utama
			var data_temp = [3]string{input_id, input_nama, fmt.Sprintf("%.1f", input_rating)}
			data = append(data, data_temp)
			fmt.Println("Data berhasil dimasukkan")

		case 2:
			var (
				input_id    string
				index       int = -1
				kiri, kanan [][3]string
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
				for i := 0; i < len(data); i++ {
					fmt.Println("ID :", data[i][0])
					fmt.Println("	Nama :", data[i][1])
					fmt.Println("	Score :", data[i][2])
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
					fmt.Println("ID :", data[i][0])
					fmt.Println("	Nama :", data[i][1])
					fmt.Println("	Score :", data[i][2])
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
				data_copy := make([][3]string, len(data))
				copy(data_copy, data)
				fmt.Println(data_copy)

				// sort data_copy by rating using ParseFloat
				for i := 0; i < len(data_copy)-1; i++ {
					idx_nilai_terbesar = i
					for j := i + 1; j < len(data_copy); j++ {
						num1, err := strconv.ParseFloat(data_copy[i][2], 2)
						if err != nil {
							fmt.Println("Error when converting to float of", data_copy[i][2])
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

				for i := 0; i < len(data_copy); i++ {
					fmt.Println("ID :", data_copy[i][0])
					fmt.Println("	Nama :", data_copy[i][1])
					fmt.Println("	Score :", data_copy[i][2])
				}

			} else {
				// Kalau misalkan datanya kosong
				fmt.Println("Data kosong")
			}
		case 6:
			var counter int = 0
			if len(data) > 0 {
				fmt.Println("Menampilkan game yang memiliki score >4.0")
				for i := 0; i < len(data); i++ {
					nilai, err := strconv.ParseFloat(data[i][2], 2)
					if err != nil {
						fmt.Println("Error when converting to float of", data[i][2])
					}
					if nilai > 4.0 {
						fmt.Println("ID :", data[i][0])
						fmt.Println("	Nama :", data[i][1])
						fmt.Println("	Score :", data[i][2])
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

	}

}
