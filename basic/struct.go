package main

import "fmt"

func main() {
	type person struct {
		name string
		age int
	}

	type student struct {
		grade int
		person
	}

	s1 := student{}
	s1.name = "Muhamad Habibi Azmi"
	s1.age = 26
	s1.grade = 2

	fmt.Println("Debug: ", s1.name)
}