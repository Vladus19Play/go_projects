package main

import (
	"fmt"
	"strconv"
)

func is_arm(n string) string {
	l := []int{}
	for i := 0; i < len(n); i++ {
		digit, _ := strconv.Atoi(string(n[i]))
		l = append(l, digit*digit*digit)
	}
	sum := 0
	for _, value := range l {
		sum += value
	}
	if sum == atoi(n) {
		return "число Армстронга"
	}
	return "не число Армстронга"
}

func atoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func main() {
	var n string
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Число %s - %s", n, is_arm(n))
}
