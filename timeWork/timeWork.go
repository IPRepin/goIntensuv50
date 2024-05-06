package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func runTiming() {
	scanner := bufio.NewScanner(os.Stdin)
	var timeLst []string
	for {
		fmt.Println("Введите время пробега 10 км: ")
		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
			}
			break
		}
		input := scanner.Text()
		if input == "" {
			break
		}
		timeLst = append(timeLst, input)
	}

	digits := make([]float64, len(timeLst))

	for i := range timeLst {
		num, err := strconv.ParseFloat(timeLst[i], 64)
		if err != nil {
			fmt.Println("Ошибка парсинга:", err)
		}
		digits[i] = num
	}

	sum := 0.0
	for _, num := range digits {
		sum += num
	}

	avrTime := sum / float64(len(timeLst))
	fmt.Printf("Средний показатель %.2f, более %d пробежек!\n", avrTime, len(timeLst))
}

func main() {
	runTiming()
}
