package main

import (
	"fmt"
)

// do a gorutine at hospital attention
func message(text string, c chan<- string) {
	c <- text
}
func multipleChannelFunction() {
	c := make(chan string, 2) // channel length reciver
	c <- "Message one"
	c <- "Message two"
	fmt.Println(len(c), cap(c))

	//Range and Close
	close(c)
	// c <- "Message three"
	for message := range c {
		fmt.Println(message)
	}

	// Select
	emailOne := make(chan string)
	emailTwo := make(chan string)
	go message("Message 1: ", emailOne)
	go message("Message 2: ", emailTwo)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-emailOne:
			fmt.Println("Email reciver from EmailOne", m1)
		case m2 := <-emailTwo:
			fmt.Println("Email reciver from EmailTwo", m2)
		}
	}
}
