package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhamad"
	names[1] = "Habibi"
	names[2] = "Azmi"

	fmt.Println(names[0]);
	fmt.Println(names[1]);
	fmt.Println(names[2]);

	var values = [3]int{
		80,
		90,
		100,
	}

	fmt.Println(values)
}