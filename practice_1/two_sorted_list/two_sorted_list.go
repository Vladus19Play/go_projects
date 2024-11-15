package main

import "fmt"

func sorted_lists(arr1 []int, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))

	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}

	return result
}

func main() {
	var n1 int
	fmt.Print("Введите длину первого массива: ")
	fmt.Scan(&n1)
	l1 := make([]int, 0, n1)
	fmt.Print("Введите элементы первого массива: ")
	for i := 0; i < n1; i++ {
		var num int
		fmt.Scan(&num)
		l1 = append(l1, num)
	}
	var n2 int
	fmt.Print("Введите длину второго массива: ")
	fmt.Scan(&n2)
	l2 := make([]int, 0, n2)
	fmt.Print("Введите элементы второго массива: ")
	for i := 0; i < n2; i++ {
		var num int
		fmt.Scan(&num)
		l2 = append(l2, num)
	}
	result := sorted_lists(l1, l2)
	fmt.Println(result)
}
