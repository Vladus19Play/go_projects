package main

import (
	"fmt"
	"math/rand"
)

func random_num(n int) {
	var rand_num = rand.Intn(100)
	var attmps int = 10
	for true {
		if attmps == 0 {
			fmt.Print("Ты проиграл!\n")
			fmt.Printf("Загаданное число %d", rand_num)
			return
		}
		if rand_num == n {
			fmt.Printf("Верно! (осталось попыток: %d) ", attmps)
			return
		} else if rand_num > n {
			attmps -= 1
			fmt.Printf("Твоё число меньше загаданного! (попыток: %d) ", attmps)
			fmt.Scan(&n)
		} else if rand_num < n {
			attmps -= 1
			fmt.Printf("Твоё число больше загаданного! (попыток: %d) ", attmps)
			fmt.Scan(&n)
		}
	}
}

func main() {
	fmt.Println("Добро пожаловать в Угадайку!")
	var n int
	fmt.Println("Попробуйте угадать число от 1 до 100, которое загадал компьютер за 10 попыток!")
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	random_num(n)
}
