package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Wow!")

	time.Sleep(10 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display: ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go DisplayNumber(i)
	}

	// time.Sleep(10 * time.Second)
}