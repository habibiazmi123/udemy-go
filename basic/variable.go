package main

import "fmt"

func main() {
	var name string

	// First Method
	name = "Muhamad Habibi Azmi";
	fmt.Println(name);

	name = "NurHijjah Arigawati";
	fmt.Println(name);

	// Second Method
	var friendName = "Lutfi";
	fmt.Println(friendName);
	
	// Third Method
	age := 20;
	fmt.Println(age);

	// Fourth Method
	var (
		firstName = "Muhamad"
		lastName = "Habibi Azmi";
	)

	fmt.Println(firstName);
	fmt.Println(lastName);
}