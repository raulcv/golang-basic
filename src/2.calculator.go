package main

import "fmt"

func calculator() {
	y := 15
	x := 3
	result := y + x
	fmt.Println("result sum: ", result)
	result = y - x
	fmt.Println("result res: ", result)
	result = y * x
	fmt.Println("result mul: ", result)
	result = y / x
	fmt.Println("result div: ", result)
}
