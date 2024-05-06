package main

import (
	"fmt"
	"math"
	"strings"
)

func hexOutput() int {
	var decnum int
	hexnum := strings.TrimSpace(input("Введите шестнадцатеричное число для преобразования: "))

	for i, digit := range hexnum {
		if digit >= '0' && digit <= '9' {
			decnum += int(digit-'0') * int(math.Pow(16, float64(len(hexnum)-i-1)))
		} else if digit >= 'A' && digit <= 'F' {
			decnum += int(digit-'A'+10) * int(math.Pow(16, float64(len(hexnum)-i-1)))
		}
	}

	return decnum
}

// input читает строку ввода.
func input(prompt string) string {
	fmt.Scanln(&prompt)
	return prompt
}

func main() {
	fmt.Println(hexOutput())
}
