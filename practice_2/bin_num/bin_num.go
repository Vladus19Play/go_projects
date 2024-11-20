package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func bin(n int) string {
	var ost string = ""
	for true {
		if n == 0 {
			break
		}
		ost += strconv.Itoa(n % 2)
		n /= 2
	}
	return reverse(ost)
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Двоичное представление: %s", bin(n))
}
