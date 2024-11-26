package main

import (
	"fmt"
	"udemy-go/database"
)

func main() {
	result:=database.GetDataBase()
	fmt.Println(result)
}