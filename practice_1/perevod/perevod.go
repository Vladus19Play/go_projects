package main

import (
	"fmt"
	"math/big"
	"strings"
)

func convertNumber(inputBase int, number string, outputBase int) string {
	decimalValue := new(big.Int)
	decimalValue.SetString(number, inputBase)
	convertedNumber := decimalValue.Text(outputBase)
	if outputBase > 10 {
		convertedNumber = strings.ToUpper(convertedNumber)
	}
	return convertedNumber
}

func main() {
	var inputBase, outputBase int
	var number string

	fmt.Print("Введите основание исходной системы счисления: ")
	fmt.Scan(&inputBase)
	fmt.Printf("Введите число в системе счисления %d: ", inputBase)
	fmt.Scan(&number)
	fmt.Print("Введите основание новой системы счисления: ")
	fmt.Scan(&outputBase)

	result := convertNumber(inputBase, number, outputBase)
	fmt.Printf("%s в системе счисления %d равно %s в системе счисления %d\n", number, inputBase, result, outputBase)
}
