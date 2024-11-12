package main

import (
	"fmt"
)

func visokos_year(year int16) {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Да, это високосный год.")
	} else {
		fmt.Println("Нет, это невисокосный год.")
	}
}

func main() {
	var year int16
	fmt.Println("Введите год!")
	fmt.Scanln(&year)
	visokos_year(year)
}
