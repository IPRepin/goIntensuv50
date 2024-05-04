package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingGame() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func main() {
	fmt.Println("Добро пожаловать в игру 'Угадай число!'")
	guessNum := guessingGame()
	var userNumber int
	for {
		fmt.Scan(&userNumber)
		if guessNum == userNumber {
			fmt.Printf("Поздравляю ты угадал число %d", userNumber)
			break
		} else if guessNum > userNumber {
			fmt.Printf("%d Слишком маленькое число попробуй еще раз", userNumber)
		} else {
			fmt.Printf("%d Слишком большое число попробуй еще раз", userNumber)
		}
	}
}
