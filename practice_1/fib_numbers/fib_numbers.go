package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	a, b := 0, 1
	fmt.Println("Ряд Фибоначчи, не превышающий", n)
	for b <= n {
		a, b = b, a+b
		fmt.Print(a, " ")
	}
}
