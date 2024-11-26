package main

import "fmt"

func getFullname() (firstName, middleName, lastName string) {
	firstName="Muhamad"
	middleName="Habibi"
	lastName="Azmi"
	return
}

func main() {
	firstName, middleName, lastName := getFullname()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}