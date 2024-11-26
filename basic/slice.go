package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// array [...]string{}
	// slice []string{}

	// months[6] = "Ubah"
	// fmt.Println(months)

	// slice reference from the array
	// slice1[0] = "Ubah"
	// fmt.Println(months)
}