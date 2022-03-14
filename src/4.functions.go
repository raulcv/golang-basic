package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}
func funWithThreeArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}
func funReturnAValue(a int) int {
	return a * 2
}
func funcReturnTwoValues(a int) (c int, d string) {
	return 2 * 3, "Thrid Argument val "
}
func principal() {
	normalFunction("hello World")
	funWithThreeArguments(34, 45, "I don't know")
	value := funReturnAValue(22)
	println("Value: ", value)
	value1, value2 := funcReturnTwoValues(33)
	fmt.Println("Value one: ", value1, "Value two: ", value2)
	value3, _ := funcReturnTwoValues(25)
	fmt.Println("Value one: ", value3)
}
