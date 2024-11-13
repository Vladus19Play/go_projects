package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Print("Введите текст: ")
	fmt.Scan(&n)
	n1 := ""
	for i := len(n) - 1; i >= 0; i-- {
		n1 += string(n[i])
	}
	fmt.Println("Обратный текст:", n1)
}
