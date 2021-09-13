package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func validate(name string, age string) (bool, error) {
	if strings.TrimSpace(name) == "" {
		return true, errors.New("Name cannot be empty")
	}
	_, err := strconv.Atoi(age)
	if err != nil {
		return false, errors.New("Age must be number")
	}
	return true, nil
}

func main() {
	var name, age string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	fmt.Print("Type your age: ")
	fmt.Scanln(&age)
	if valid, err := validate(name, age); valid {
		fmt.Println("Hello", name, "your age is ", age)
	} else {
		fmt.Println(err.Error())
	}
}
