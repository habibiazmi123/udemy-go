package main

import "fmt"

func main() {
	var result = 10 + 10;
	fmt.Println(result);

	var a = 10;
	var b = 10;
	var c = a * b;
	fmt.Println(c);

	// Augmented Assignment
	var i = 10;
	i += 10;
	fmt.Println(i);

	// Unary Assignment
	i++;
	fmt.Println(i);

	var negative = -100;
	var positive = 100;

	fmt.Println(negative);
	fmt.Println(positive);
}