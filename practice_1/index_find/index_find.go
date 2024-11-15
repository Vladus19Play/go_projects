package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSubstring(input string, output string) int {
	hLen := len(input)
	nLen := len(output)

	if nLen == 0 || hLen < nLen {
		return -1
	}

	for i := 0; i <= hLen-nLen; i++ {
		found := true
		for j := 0; j < nLen; j++ {
			if input[i+j] != output[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	var output string
	fmt.Print("Введите искомую подстроку: ")
	fmt.Scan(&output)
	index := findSubstring(input, output)
	fmt.Print("Индекс совпадения: ", index)
}
