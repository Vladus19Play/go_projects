package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func palindrome(n string) string {
	total := ""
	n = strings.ToLower(n)
	for _, i := range n {
		if unicode.IsLetter(i) {
			total += string(i)
			continue
		}
	}
	if total == reverse(total) {
		return "Да!"
	} else {
		return "Нет:("
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Введите строку!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := scanner.Text()
	fmt.Println("Является ли введённая строка палиндромом? -", palindrome(n))
}
