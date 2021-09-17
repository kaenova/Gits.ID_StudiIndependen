package main

import (
	"fmt"
	"math"
)

// Defining for Persegi
type PersegiPanjnag struct {
	panjang float32
	lebar   float32
}

func (p PersegiPanjnag) Keliling() float32 {
	return (2 * p.panjang) + (2 * p.lebar)
}

func (p PersegiPanjnag) Luas() float32 {
	return p.panjang * p.lebar
}

// Defining for Segitiga Sama sisi
type Segitiga struct {
	sisi float32
	/*
		Ini kalau bisa ke inisialisasiin berdasarkan nilai alas.
		Bagaimana ya kak caranya?
		Jadi di segitiganya hanya masukkin 1 nilai saja.
	*/
}

func (s Segitiga) Tinggi() float32 {
	setengah_alas_kuadrat := math.Pow(float64(0.5*s.sisi), 2)
	alas_kuadrat := math.Pow(float64(s.sisi), 2)

	return float32(
		math.Pow(float64(alas_kuadrat-setengah_alas_kuadrat), 0.5))
}

func (s Segitiga) Keliling() float32 {
	// Belum selesai dibuat
	return s.sisi * 3
}

func (s Segitiga) Luas() float32 {
	return 0.5 * s.sisi * s.Tinggi()
}

// Making interface
type Hitung interface {
	Luas() float32
	Keliling() float32
}

func main() {
	var bangunDatar Hitung

	fmt.Printf("=== %-20s\n", "Persegi Panjang")
	bangunDatar = PersegiPanjnag{panjang: 10, lebar: 5}
	fmt.Printf("%10s : %.2f\n", "Keliling", bangunDatar.Keliling())
	fmt.Printf("%10s : %.2f\n", "Luas", bangunDatar.Luas())

	fmt.Printf("=== %-20s\n", "Segitiga Sama Sisi")
	bangunDatar = Segitiga{sisi: 10}
	fmt.Printf("%10s : %.2f\n", "Keliling", bangunDatar.Keliling())
	fmt.Printf("%10s : %.2f\n", "Luas", bangunDatar.Luas())
}
