package main

import (
	"fmt"
	"sync"
	"time"
)

func saysomething(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
func goRoutineWaitGroupFunction() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)
	go saysomething("World", &wg)
	wg.Wait()
	go func(text string) {
		fmt.Println(text)
	}("Bye!!!")
	time.Sleep(time.Second * 1)
}
