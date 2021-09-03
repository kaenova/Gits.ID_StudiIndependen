/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		nilai, input, mod          int
		final_rupiah, final_dollar string
		minus, first_conversion    bool
	)
	// Currency Initialization
	minus = false
	final_rupiah = ",00"
	first_conversion = true

	//Input
	fmt.Print("Masukkan jumlah nilai : ")
	fmt.Scanf("%d", &input)

	nilai = input

	// Minus check
	if nilai < 0 {
		minus = true
		nilai = nilai * -1
	}

	// Konversi nilai rupiah
	if nilai > 0 {

		for nilai > 0 {
			mod = nilai % 1000
			nilai = nilai / 1000

			if first_conversion {
				first_conversion = false
			} else {
				final_rupiah = "." + final_rupiah
			}

			if mod == 0 {
				final_rupiah = "000" + final_rupiah
			} else if mod < 10 {
				// Kurang dari sepuluh
				final_rupiah = strconv.Itoa(mod) + final_rupiah
				if nilai > 0 {
					// Kalau misalkan ada yang masih bisa dibagi lagi
					final_rupiah = "00" + final_rupiah
				}
			} else if mod/10 < 10 {
				// Kurang dari 100
				final_rupiah = strconv.Itoa(mod) + final_rupiah
				if nilai > 0 {
					final_rupiah = "0" + final_rupiah
				}
			} else {
				final_rupiah = strconv.Itoa(mod) + final_rupiah
			}
		}
	} else {
		final_rupiah = "0" + final_rupiah
	}

	// Finalisasi
	final_rupiah = "Rp " + final_rupiah
	if minus {
		final_rupiah = "- " + final_rupiah
	}

	// Output Rupiah
	fmt.Println("Hasil konversi Rupiah :", final_rupiah)

	// Konversi ke Dollar
	fmt.Println("Dengan nilai tukar Dolar ke Rupiah 1 USD = Rp 14.000.000,00")

	// Buat menjadi 2 precision
	var nilai_dollar, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float32(input)/14000.0), 2)

	if err == nil {
		final_dollar = ""

		// Check apakah ada koma?
		nilai_koma := int(nilai_dollar*100.0) % 100
		koma := true
		if nilai_koma == 0 {
			koma = false
		}

		if koma {
			if nilai_koma < 10 {
				final_dollar = ",0" + strconv.Itoa(nilai_koma)
			} else {
				final_dollar = "," + strconv.Itoa(nilai_koma)
			}
		} else {
			final_dollar = ",00"
		}

		var selain_koma_dollar int = int(nilai_dollar)

		first_conversion = true

		//Konversi nilai selain koma ke dollar
		if selain_koma_dollar > 0 {

			for selain_koma_dollar > 0 {
				mod = selain_koma_dollar % 1000
				selain_koma_dollar = selain_koma_dollar / 1000

				if first_conversion {
					first_conversion = false
				} else {
					final_dollar = "." + final_dollar
				}

				if mod == 0 {
					final_dollar = "000" + final_dollar
				} else if mod < 10 {
					// Kurang dari sepuluh
					final_dollar = strconv.Itoa(mod) + final_dollar
					if selain_koma_dollar > 0 {
						// Kalau misalkan ada yang masih bisa dibagi lagi
						final_dollar = "00" + final_dollar
					}
				} else if mod/10 < 10 {
					// Kurang dari 100
					final_dollar = strconv.Itoa(mod) + final_dollar
					if selain_koma_dollar > 0 {
						final_dollar = "0" + final_dollar
					}
				} else {
					final_dollar = strconv.Itoa(mod) + final_dollar
				}
			}
		} else {
			final_dollar = "0" + final_dollar
		}

		// Finalisasi
		final_dollar = "$ " + final_dollar
		if minus {
			final_dollar = "- " + final_dollar
		}

		// Output
		fmt.Println("Hasil konversi Dollar :", final_dollar)

	} else {
		// Panic saat mengubah menjadi 2 decimal placess
		fmt.Println("Error!")
	}

}
