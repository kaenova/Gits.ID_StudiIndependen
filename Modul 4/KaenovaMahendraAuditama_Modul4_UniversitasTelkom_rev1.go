package main

import (
	"fmt"
	"strconv"
)

func main() {

	var pilihan int
	const (
		// Based on Google (06/09/20212)
		rupiah_ke_dollar float32 = 0.00007
		rupiah_ke_euro   float32 = 0.000059
	)

	/*
		================
		Pilih menu penukaran
		1. Rupiah ke Dollar
		2. Rupiah ke Euro
		3. GBP ke Knut
		Masukkan Nomor :
		===============
		Masukkan nominal :
		Hasil Konversi :
	*/
	fmt.Println("================")
	fmt.Println("Pilih menu penukaran")
	fmt.Println("1. Rupiah ke Dollar")
	fmt.Println("2. Rupiah ke Euro")
	fmt.Println("3. GBP ke Knut")
	fmt.Print("Masukkan Nomor : ")
	fmt.Scanf("%d", &pilihan)
	fmt.Println("================")

	if pilihan == 1 {
		var (
			rupiah float32
			dollar float32
		)

		fmt.Print("Masukkan nominal (Rupiah) : ")
		fmt.Scan(&rupiah)
		dollar = rupiah * rupiah_ke_dollar
		fmt.Printf("Hasil Konversi (Dollar) : %s", DollarEuroFormatting(dollar, "dollar"))

	} else if pilihan == 2 {
		var (
			rupiah float32
			euro   float32
		)

		fmt.Print("Masukkan nominal (Rupiah) : ")
		fmt.Scan(&rupiah)
		euro = rupiah * rupiah_ke_euro
		fmt.Printf("Hasil Konversi (Euro) : %s", DollarEuroFormatting(euro, "euro"))

	} else if pilihan == 3 {
		/*
			- 1 GBP = 100 Knuts
			- 1 Galleon = 17 Sickle
			- 1 Sickle = 29 Knut

			Perhiungan Manual:
			1 galleon = 17 * 29 = 493 knut
		*/
		var (
			gbp, knut, sickle, galleon int = 0, 0, 0, 0
		)

		fmt.Print("Masukkan nominal (GBP) : ")
		fmt.Scan(&gbp)

		// Konversi gbp ke knut
		knut = gbp * 100
		fmt.Println("Jumlah knut yang di dapat :", knut, " Knut (s)")

		// Penukaran dari knut ke galleon
		galleon = knut / 493
		sisa_knut := knut % 493
		fmt.Println("Hasil penukaran mendapatkan :", galleon, " Galleon (s)")

		// Sisa knut dari penukaran galleon ke sickle
		sickle = sisa_knut / 29
		sisa_sickle := sisa_knut % 29
		fmt.Println("Sisa ditukar menjadi :", sickle, " Sickle (s)")

		fmt.Println("Keping knut yang tersisa :", sisa_sickle, " Knut (s)")

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func DollarEuroFormatting(nominal float32, currency string) string {
	var final string
	var minus bool

	if nominal < 0 {
		minus = true
	} else {
		minus = false
	}

	prec, err := strconv.ParseFloat(fmt.Sprintf("%.2f", nominal), 2)

	if err == nil {
		final = ""

		// Check apakah ada koma?
		nilai_koma := int(prec*100.0) % 100
		koma := true
		if nilai_koma == 0 {
			koma = false
		}

		if koma {
			if nilai_koma < 10 {
				final = ",0" + strconv.Itoa(nilai_koma)
			} else {
				final = "," + strconv.Itoa(nilai_koma)
			}
		} else {
			final = ",00"
		}

		var selain_koma int = int(prec)

		first_conversion := true

		//Konversi nilai selain koma ke dollar
		if selain_koma > 0 {

			for selain_koma > 0 {
				mod := selain_koma % 1000
				selain_koma = selain_koma / 1000

				if first_conversion {
					first_conversion = false
				} else {
					final = "." + final
				}

				if mod == 0 {
					final = "000" + final
				} else if mod < 10 {
					// Kurang dari sepuluh
					final = strconv.Itoa(mod) + final
					if selain_koma > 0 {
						// Kalau misalkan ada yang masih bisa dibagi lagi
						final = "00" + final
					}
				} else if mod/10 < 10 {
					// Kurang dari 100
					final = strconv.Itoa(mod) + final
					if selain_koma > 0 {
						final = "0" + final
					}
				} else {
					final = strconv.Itoa(mod) + final
				}
			}
		} else {
			final = "0" + final
		}

		// Finalisasi
		if currency == "dollar" {
			final = "$ " + final
		} else if currency == "euro" {
			final = "€ " + final
		}

		if minus {
			final = "- " + final
		}

		// Output
		return final

	} else {
		// Panic saat mengubah menjadi 2 decimal placess
		return "err"
	}
}
