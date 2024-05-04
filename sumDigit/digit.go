package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumDigit() int {
	var digitStr string
	fmt.Print("Введите числа через запятую:")
	fmt.Scan(&digitStr)
	elements := strings.Split(digitStr, ",")
	digits := make([]int, len(elements))
	for i, element := range elements {
		digits[i], _ = strconv.Atoi(element)
	}
	total := 0
	for _, digit := range digits {
		total += digit
	}
	return total
}

func main() {
	fmt.Println(sumDigit())
}
