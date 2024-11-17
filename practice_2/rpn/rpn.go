package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rpn(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		}
	}
	return stack[0]
}

func main() {
	fmt.Println("Введите строку по обратной польской записи (все отдельные символы через пробел): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	tokens := strings.Fields(input)
	fmt.Println(rpn(tokens))
}
