package main

import "fmt"

type Res struct {
	Angka int
	Kata  string
}

func (r Res) Init() {
	r.Angka = 5
	r.Kata = "Bruh"
}

func main() {
	var bruh Res
	bruh.Init()
	fmt.Println(bruh.Angka, bruh.Kata)
}
