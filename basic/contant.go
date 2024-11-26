package main

import "fmt"

func main() {
	const firstName string = "Muhamad"
	const lastName = "Habibi Azmi"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		age = 10
		address = "Jl Antapani"
	)

	fmt.Println(age)
	fmt.Println(address)
}