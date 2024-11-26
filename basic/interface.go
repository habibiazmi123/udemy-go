package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2;
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6;
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12;
}

func main() {
	// var bangunDatar hitung;

	// bangunDatar = persegi{10}
	// fmt.Println("Luas: ", bangunDatar.luas())
	// fmt.Println("Keliling: ", bangunDatar.keliling())

	// fmt.Println("--------------------------------")

	// bangunDatar = lingkaran{14}
	// fmt.Println("Luar: ", bangunDatar.luas())
	// fmt.Println("Keliling: ", bangunDatar.keliling())
	// fmt.Println("Jari: ", bangunDatar.(lingkaran).jariJari())

	fmt.Println("--------------------------------")

	var bangunRuang hitung = &kubus{14}
	fmt.Println("luas      :", bangunRuang.luas())
    fmt.Println("keliling  :", bangunRuang.keliling())
    fmt.Println("volume    :", bangunRuang.volume())

}