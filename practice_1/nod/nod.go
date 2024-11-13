package main

import (
	"fmt"
)

func nod(a, b int) int {
	mx := a
	mn := b
	if a < b {
		mx, mn = b, a
	}
	if a+b == b || a+b == a {
		return a + b
	}
	for {
		if mx%mn == 0 {
			return mn
		}
		ost := mx % mn
		mx = mn
		mn = ost
	}
}

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Println("НОД:", nod(a, b))
}
