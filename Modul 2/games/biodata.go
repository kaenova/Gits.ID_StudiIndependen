package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var (
		umur int
	)

	in := bufio.NewReader(os.Stdin)

	fmt.Print("Write your name : ")
	nama, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Write your university : ")
	universitas, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Write your age : ")
	fmt.Scanf("%d", &umur)

	fmt.Print("Write your hobby : ")
	hobi, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	nama = strings.ReplaceAll(nama, "\n", "")
	universitas = strings.ReplaceAll(universitas, "\n", "")
	hobi = strings.ReplaceAll(hobi, "\n", "")

	fmt.Printf("my name is %s from %s my age %d and my hobby %s.", nama, universitas, umur, hobi)	
}