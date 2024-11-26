package main

import "fmt"

func main() {
	name:="Cumi"

	if name == "Azmi" {
		fmt.Println("Hai Azmi")
	} else if name == "Cumi" {
		fmt.Println("Hai Cumi")
	} else {
		fmt.Println("Hai, Boleh kenalan?")
	}

	if length := len(name); length > 3 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Oke Sesuai")
	}
}