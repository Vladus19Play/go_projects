package main

import (
	"fmt"
)

func calc() {
	fmt.Println("Введите первое число: ")
	var a int
	fmt.Scanln(&a)
	fmt.Println("Введите арифмитический знак: (+, -, *, /, ^, %) ")
	var znak string
	fmt.Scanln(&znak)
	fmt.Println("Введите второе число: ")
	var b int
	fmt.Scanln(&b)
	switch znak {
	case "-":
		fmt.Println("Ваш ответ:", a-b)
	case "+":
		fmt.Println("Ваш ответ:", a+b)
	case "*":
		fmt.Println("Ваш ответ:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль невозможно!")
			break
		}
		a := float32(a)
		b := float32(b)
		fmt.Println("Ваш ответ:", a/b)
	case "%":
		fmt.Println("Ваш ответ:", a%b)
	case "^":
		fmt.Println("Ваш ответ:", a-b)
	default:
		fmt.Println("Введите доступную операцию.")
		calc()
	}
	fmt.Println("Посчитать еще? (y/n): ")
	var ans string
	fmt.Scanln(&ans)
	if ans == "y" {
		calc()
	} else {
		return
	}
}

func main() {
	fmt.Println("Добро пожаловать в калькулятор!")
	calc()
}
