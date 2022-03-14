package main

import (
	"fmt"
)

func keyWordsFunction() {
	//Defer
	defer fmt.Println("hola")
	fmt.Println("Mundo")
	//Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println("I value: ", i)
		//Continue
		if i == 2 {
			fmt.Println("I value is Two")
			continue
		}
		//break
		if i == 8 {
			fmt.Println("I value is Eight")
			break
		}
	}
}
