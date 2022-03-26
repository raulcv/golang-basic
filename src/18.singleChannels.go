package main

import (
	"fmt"
)

func say(text string, c chan<- string) {
	c <- text
}
func singleChannels() {
	c := make(chan string, 1)
	fmt.Println("Hello, world!")
	go say("Bye", c)
	fmt.Println(<-c)
}
