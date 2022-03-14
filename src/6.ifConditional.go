package main

import (
	"fmt"
	"log"
	"strconv"
)

func ifConditionalFunction() {
	value1 := 1
	value2 := 3
	if value1 == 1 {
		fmt.Println("It's one")
	} else {
		fmt.Println("It's not one")
	}
	if value1 == 1 && value2 == 2 {
		fmt.Println("The values are right!!!")
	} else if value1 == 1 || value2 == 5 {
		fmt.Println("One of them values is right!!!")
	} else {
		fmt.Println("The values are not right!!!")
	}

	value, err := strconv.Atoi("45")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value of string convet to number: ", value)
}
