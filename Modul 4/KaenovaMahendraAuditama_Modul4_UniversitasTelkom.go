package main

import "fmt"

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
		dollar = rupiah / rupiah_ke_dollar
		fmt.Printf("Hasil Konversi (Dollar) : %.2f", dollar)

	} else if pilihan == 2 {
		var (
			rupiah float32
			euro   float32
		)

		fmt.Print("Masukkan nominal (Rupiah) : ")
		fmt.Scan(&rupiah)
		euro = rupiah / rupiah_ke_euro
		fmt.Printf("Hasil Konversi (Euro) : %.2f", euro)

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
