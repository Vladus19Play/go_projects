package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func the_longest(l string) string {
	l1 := strings.Fields(l)
	for i, word := range l1 {
		l1[i] = strings.Trim(word, ",.!?'-()@#$%^&*=+{}[]><")
	}
	maxWord := ""
	for _, word := range l1 {
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	return maxWord
}

func main() {
	fmt.Println("Введите предложение!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	fmt.Println("Самое длинное слово:", the_longest(l))
}
