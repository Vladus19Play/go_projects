package main

import (
	"fmt"
)

func table(a int, b int) {
	fmt.Println("\nТаблица умножения: ")
	for i := a; i <= b; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	var a, b int
	fmt.Print("Введите левую границу: ")
	fmt.Scan(&a)
	fmt.Print("Введите правую границу: ")
	fmt.Scan(&b)
	table(a, b)
}
