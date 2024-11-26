package main

import "fmt"

func getFullname() (string, string, string) {
	return "Muhamad", "Habibi", "Azmi"
}

func main() {
	firstName, middleName, _ := getFullname()
	fmt.Println(firstName)
	fmt.Println(middleName)
	// fmt.Println(lastName)
}