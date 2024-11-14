package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Print("Введите длину массива: ")
	fmt.Scan(&n)
	l := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	sort.Slice(l, func(i, j int) bool {
		return math.Abs(float64(l[i])) < math.Abs(float64(l[j]))
	})

	fmt.Println(l)
}
