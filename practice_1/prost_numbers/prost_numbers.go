package main

import (
	"fmt"
)

func index_of(n int, data []int) int {
	for k, v := range data {
		if n == v {
			return k
		}
	}
	return -1
}

func reshetoSundarama(m int, n int) []int {
	allNumbers := make([]int, 0, (n-1)/2)
	for i := 1; i <= (n-1)/2; i++ {
		allNumbers = append(allNumbers, i)
	}

	for i := 1; i < (n - 1); i++ {
		for j := 1; j < (n - 1); j++ {
			num := 2*i*j + i + j
			for k, v := range allNumbers {
				if v == num {
					allNumbers = append(allNumbers[:k], allNumbers[k+1:]...)
					break
				}
			}
		}
	}

	for i := range allNumbers {
		allNumbers[i] = allNumbers[i]*2 + 1
	}

	for i, j := 0, len(allNumbers)-1; i < j; i, j = i+1, j-1 {
		allNumbers[i], allNumbers[j] = allNumbers[j], allNumbers[i]
	}
	allNumbers = append(allNumbers, 2)
	allNumbers = append(allNumbers, 1)
	for i, j := 0, len(allNumbers)-1; i < j; i, j = i+1, j-1 {
		allNumbers[i], allNumbers[j] = allNumbers[j], allNumbers[i]
	}

	for _, i := range allNumbers {
		if i == m {
			return allNumbers[index_of(i, allNumbers):]
		} else if i > m {
			return allNumbers[index_of(i, allNumbers):]
		}
	}
	return allNumbers
}

func main() {
	var m, n int
	fmt.Print("Введите начало диапазона: ")
	fmt.Scan(&m)
	fmt.Print("Введите конец диапазона: ")
	fmt.Scan(&n)

	fmt.Printf("Все простые числа от %d до %d = %v\n", m, n, reshetoSundarama(m, n))
}
