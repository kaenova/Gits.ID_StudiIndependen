package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		base_pangkat float64
		pengali      float64
		pengulang    float64
		ulang        bool = true
	)

	for ulang {
		fmt.Print("Masukkan base : ")
		fmt.Scanln(&base_pangkat)
		fmt.Print("Masukkan pengali : ")
		fmt.Scanln(&pengali)
		fmt.Print("Maximum pangkat berapa ? ")
		fmt.Scanln(&pengulang)
		fmt.Print("Sum = ")
		Pangkat(0, pengali, pengulang+1, 0, base_pangkat)
	}

}

func Pangkat(jumlah, pengali, pangkat_max, pangkat_now, base float64) {
	if pangkat_now < pangkat_max {
		temp := pengali * (math.Pow(base, pangkat_now))
		jumlah = jumlah + temp
		if pangkat_now == pangkat_max-1 {
			fmt.Printf(" %.2f =", temp)
		} else {
			fmt.Printf(" %.2f +", temp)
		}
		Pangkat(jumlah, pengali, pangkat_max, pangkat_now+1, base)
	} else {
		fmt.Printf(" %.2f\n", jumlah)
	}
}
