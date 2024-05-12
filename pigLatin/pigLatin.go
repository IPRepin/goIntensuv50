package main

import (
	"bufio"
	"os"
	"strings"
)

func pigLatin() string {
	vowels := []rune("aeiouAEIOU") // Include uppercase vowels as well
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if strings.ContainsRune(string(vowels), rune(text[0])) {
		return text + "way"
	} else {
		return text[1:] + string(text[0]) + "ay"
	}
}

func main() {
	println(pigLatin())
}
