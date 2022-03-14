package main

import "fmt"

func loopFunction() {
	//for
	for i := 0; i <= 10; i++ {
		fmt.Println("Value for i: ", i)
	}
	fmt.Printf("\n")
	//For While
	counter := 0
	for counter < 10 {
		fmt.Println("Value for counter: ", counter)
		counter++
	}
	//For forever loop
	/*
		counterForever := 0
		for {
			fmt.Println("Value for counterForever ", counterForever)
			counterForever++
		}
	*/
}
