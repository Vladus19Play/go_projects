package main

import (
	"fmt"
)

func recur_fib(n int) int {
	if n <= 1 {
		return n
	}
	return recur_fib(n-1) + recur_fib(n-2)
}

func main() {
	var n int
	fmt.Print("Введите число членов последовательности: ")
	fmt.Scan(&n)
	fmt.Printf("Первые %d чисел Фибоначчи:\n", n)
	for i := 0; i < n; i++ {
		fmt.Print(recur_fib(i))
		fmt.Print(" ")
	}
}
