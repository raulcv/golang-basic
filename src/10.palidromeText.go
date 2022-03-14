package main

import (
	"fmt"
	"strings"
)

func isPalidrome(text string) {
	text = strings.ToUpper(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("It's Palidrome")
	} else {
		fmt.Println("It's not Palidrome")
	}

}
func palidromeFunction() {
	slice := []string{"hello", "what's", "happen"}
	for _, value := range slice {
		fmt.Println(value)
	}
	for i, value := range slice {
		fmt.Println(i, value)
	}
	for i := range slice {
		fmt.Printf("%d: ", i)
		fmt.Println(slice[i])
	}

	value := "Amor a roMa"
	isPalidrome(value)
}
