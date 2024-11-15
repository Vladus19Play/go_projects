package main

import (
	"fmt"
)

func TwoSegments(a1, b1, a2, b2 float64) bool {
	return maxFloat64(a1, a2) <= minFloat64(b1, b2)
}

func minFloat64(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func maxFloat64(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func ThreeSegments(a1, b1, a2, b2, a3, b3 float64) bool {
	return TwoSegments(a1, b1, a2, b2) &&
		TwoSegments(a2, b2, a3, b3) &&
		TwoSegments(a1, b1, a3, b3)
}

func main() {
	var a1, b1, a2, b2, a3, b3 float64
	fmt.Print("Введите начало и конец 1 отрезка: ")
	fmt.Scan(&a1)
	fmt.Scan(&b1)
	fmt.Print("Введите начало и конец 2 отрезка: ")
	fmt.Scan(&a2)
	fmt.Scan(&b2)
	fmt.Print("Введите начало и конец 3 отрезка: ")
	fmt.Scan(&a3)
	fmt.Scan(&b3)
	if ThreeSegments(a1, b1, a2, b2, a3, b3) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
