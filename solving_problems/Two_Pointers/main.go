package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	characters := []rune(s)
	left := 0
	right := len(characters) - 1
	for left < right {

		temp := characters[left]
		characters[left] = characters[right]
		characters[right] = temp
		left++
		right--
	}

	return string(characters)
}
func main() {
	right := 0
	left := 0
	word := "Gileno Cordeiro"
	var response []string
	for right < len(word) {
		characters := []rune(word)
		for _, ch := range characters {
			if ch != ' ' {
				right++
			} else {
				response = append(response, Reverse(word[left:right]))
				left = right
				right++

			}

		}
	}
	response = append(response, Reverse(word[left:right]))

	fmt.Println(strings.Join(response, " "))

	fmt.Println(Reverse(word))

}
