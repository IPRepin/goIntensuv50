package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumDigit() int {
	var digit_str string
	fmt.Print("Введите числа через запятую:")
	fmt.Scan(&digit_str)
	elements := strings.Split(digit_str, ",")
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
