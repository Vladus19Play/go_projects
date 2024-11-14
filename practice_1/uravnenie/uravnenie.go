package main

import (
	"fmt"
	"math"
)

func quadraticRoots(a, b, c float64) (complex128, complex128) {
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	} else if discriminant == 0 {
		root := -b / (2 * a)
		return complex(root, 0), complex(root, 0)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
		return complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)
	}
}

func main() {
	var a, b, c float64
	fmt.Print("Введите параметр 'a': ")
	fmt.Scan(&a)
	fmt.Print("Введите параметр 'b': ")
	fmt.Scan(&b)
	fmt.Print("Введите параметр 'с': ")
	fmt.Scan(&c)

	root1, root2 := quadraticRoots(a, b, c)
	fmt.Println("Корни уравнения:", root1, "и", root2)
}
