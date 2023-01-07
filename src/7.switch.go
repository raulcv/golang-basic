package main

import "fmt"

func switchFunction() {
	switch module := 5 % 2; module {
	case 0:
		fmt.Println("It's perfect!")
	default:
		fmt.Println("It's not perfect!")
	}
	//without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("It's Okay!")
	case value < 0:
		fmt.Println("It's highter than 1!")
	default:
		fmt.Println("It's not Okay!")
	}
}
