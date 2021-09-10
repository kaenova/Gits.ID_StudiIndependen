/*
Kaenova Mahendra Auditama
Universitas Telkom
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var rek_max int
	var pilihan string
	var ulang bool = true
	var pengali float32

	for ulang {
		fmt.Println("===================")
		fmt.Println("Menu")
		fmt.Println("a. Sum = 20 + 10 + 6.66 + 5 + ....")
		fmt.Println("b. Sum = 100 + 50 + 33.33 + 25 + ....")
		fmt.Println("c. Sum = 1 + 0.5 + 0.33 + 0.25 + ....")
		fmt.Println("d. Sum = 1/2x + 2/4x + 3/9x + 4/16x + ....")
		fmt.Println("e. Sum = 2x/1 + 4x/2 + 9x/3 + 16x/4 + ....")
		fmt.Println("f. keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilihan)
		fmt.Println("===================")

		switch pilihan {
		case "a":
			fmt.Print("Masukkan panjang rekursif : ")
			fmt.Scanln(&rek_max)
			fmt.Print("Sum = ")
			RekursifAwal(20, 0, 0, rek_max)
		case "b":
			fmt.Print("Masukkan panjang rekursif : ")
			fmt.Scanln(&rek_max)
			fmt.Print("Sum = ")
			RekursifAwal(100, 0, 0, rek_max)
		case "c":
			fmt.Print("Masukkan panjang rekursif : ")
			fmt.Scanln(&rek_max)
			fmt.Print("Sum = ")
			RekursifAwal(1, 0, 0, rek_max)
		case "d":
			fmt.Print("Masukkan panjang rekursif : ")
			fmt.Scanln(&rek_max)
			fmt.Print("Masukkan pengali (x) : ")
			fmt.Scanln(&pengali)
			fmt.Print("Sum = ")
			RekursifD(0, pengali, 0, rek_max)
		case "e":
			fmt.Print("Masukkan panjang rekursif : ")
			fmt.Scanln(&rek_max)
			fmt.Print("Masukkan pengali (x) : ")
			fmt.Scanln(&pengali)
			fmt.Print("Sum = ")
			RekursifE(0, pengali, 0, rek_max)
		case "f":
			ulang = false
		default:
			fmt.Println("Pilihan tidak dikenal")
		}
		rek_max = 0
		pilihan = ""
	}
}

func RekursifAwal(awal, jumlah float32, current, rek_max int) {
	if current < rek_max {
		temp := awal / float32(current+1)
		if current == rek_max-1 {
			fmt.Printf(" %.2f =", temp)
		} else {
			fmt.Printf(" %.2f +", temp)
		}
		jumlah = jumlah + temp
		RekursifAwal(awal, jumlah, current+1, rek_max)
	} else {
		fmt.Printf(" %.2f\n", jumlah)
	}
}

func RekursifD(jumlah, pengali float32, current, rek_max int) {
	var penyebut float64
	if current < rek_max {
		pembilang := current + 1
		if current == 0 {
			penyebut = 2
		} else {
			penyebut = math.Pow(float64(pembilang), 2.0) * float64(pengali)
		}
		if current == rek_max-1 {
			fmt.Printf(" %d/%.2f =", pembilang, penyebut)
		} else {
			fmt.Printf(" %d/%.2f +", pembilang, penyebut)
		}
		jumlah = (float32(pembilang) / float32(penyebut)) + jumlah
		RekursifD(jumlah, pengali, current+1, rek_max)
	} else {
		fmt.Printf(" %.2f\n", jumlah)
	}
}

func RekursifE(jumlah, pengali float32, current, rek_max int) {
	var pembilang float64
	if current < rek_max {
		penyebut := current + 1
		if current == 0 {
			pembilang = 2
		} else {
			pembilang = math.Pow(float64(penyebut), 2.0) * float64(pengali)
		}
		if current == rek_max-1 {
			fmt.Printf(" %.2f/%d =", pembilang, penyebut)
		} else {
			fmt.Printf(" %.2f/%d +", pembilang, penyebut)
		}
		jumlah = (float32(pembilang) / float32(penyebut)) + jumlah
		RekursifE(jumlah, pengali, current+1, rek_max)
	} else {
		fmt.Printf(" %.2f\n", jumlah)
	}
}
