package main

import (
	"fmt"
	"math"
)

func find_s(v float64, a float64, t float64) {
	var s float64
	s = v*t + (a*math.Pow(t, 2.0))/2
	fmt.Print(s, " м")
}

func main() {
	var v, a, t float64
	fmt.Print("Введите скорость (м/c): ")
	fmt.Scan(&v)
	fmt.Print("Введите ускорение (м/с^2): ")
	fmt.Scan(&a)
	fmt.Print("Введите время (с): ")
	fmt.Scan(&t)
	fmt.Print("S = ")
	find_s(v, a, t)
}
