package main

import "fmt"

func main() {
	persons := map[string]string{
		"name": "Muhamad habibi Azmi",
		"address": "Antapani",
	}

	persons["title"] = "Fullstack Development"

	fmt.Println(len(persons))
	fmt.Println(persons)
	fmt.Println(persons["name"])
	fmt.Println(persons["address"])
	fmt.Println(persons["title"])

	var books = make(map[string]string)
	books["title"] = "Book Programming"
	books["author"] = "MRx"
	fmt.Println(books)

	delete(books, "author")

	fmt.Println(books)
}