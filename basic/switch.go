package main

import "fmt"

func main() {
	name:="Azmi"

	switch name {
		case "Azmi":
			fmt.Println("Hai Azmi")
		case "Cumi":
			fmt.Println("Hai Cumi")
		default:
			fmt.Println("Hai, Boleh kenalan?")	
	}
}