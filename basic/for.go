package main

import "fmt"

func main() {
	count := 1;
	for count <= 10 {
		fmt.Println("Perulangan ke", count)
		count++;
	}

	slice := []string{"Ujang", "Joko", "Wahyu"}
	for i := 0; i < len(slice); i++ {
		fmt.Println("Nama Saya", slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, value);
	}

	persons := make(map[string]string)
	persons["title"] = "Azmi"
	persons["description"] = "Muhamad Habibi Azmi"

	for _, value := range persons {
		fmt.Println(value)
	}
}