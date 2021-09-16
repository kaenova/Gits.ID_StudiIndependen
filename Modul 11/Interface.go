package main

import (
	"fmt"
	"math"
)

// Defining for Persegi
type Persegi struct {
	panjang float32
	lebar   float32
}

func (p Persegi) Keliling() float32 {
	return (2 * p.panjang) + (2 * p.lebar)
}

func (p Persegi) Luas() float32 {
	return p.panjang * p.lebar
}

// Defining for Segitiga Sama sisi
type Segitiga struct {
	alas   float32
	tinggi float32
	/*
		Ini kalau bisa ke inisialisasiin berdasarkan nilai alas. Bagaimana ya kak caranya?
		Jadi di segitiganya hanya masukkin 1 nilai saja.
	*/
}

func (s Segitiga) init() {
	// Ini hanya percobaan
	setengah_alas_kuadrat := math.Pow(float64(0.5*s.alas), 2)
	alas_kuadrat := math.Pow(float64(0.5*s.alas), 2)
	s.tinggi = float32(
		math.Pow((float64(alas_kuadrat - setengah_alas_kuadrat)), 0.5))
}

func (s Segitiga) Keliling() float32 {
	// Belum selesai dibuat
	return 0
}

func (s Segitiga) Luas() float32 {
	return 0.5 * s.alas * s.tinggi
}

// Making interface
type Hitung interface {
	Luas() float32
	Keliling() float32
}

func main() {
	s := Segitiga{1, 0}
	fmt.Println(s.tinggi)
}
