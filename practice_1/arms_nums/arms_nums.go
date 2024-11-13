package main

import (
	"fmt"
	"math"
	"strconv"
)

func arm(a, b int) []int {
	l := []int{}
	for i := a; i <= b; i++ {
		l1 := len(strconv.Itoa(i))
		num := 0
		for _, ch := range strconv.Itoa(i) {
			digit, _ := strconv.Atoi(string(ch))
			num += int(math.Pow(float64(digit), float64(l1)))
		}
		if num == i {
			l = append(l, i)
		}
	}
	return l
}

func main() {
	var a, b int
	fmt.Print("Введите начало диапазона: ")
	fmt.Scan(&a)
	fmt.Print("Введите конец диапазона: ")
	fmt.Scan(&b)
	fmt.Println("Числа Армстронга на заданном диапазоне:", arm(a, b))
}
