package main

import (
	"fmt"
	"math"
)

func is_prost(n int) string {
	for i := 2; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return "Составное"
		}
	}
	return "Простое"
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Число %d - %s", n, is_prost(n))
}
