package main

import "fmt"

func fmtLogs() {
	//Declare variables
	helloMessage := "Hello"
	worldMessage := "World"
	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)
	//Printtf
	name := "Raul CV"
	cantAnimals := 14
	city := "Lima - Peru"
	fmt.Printf("%s has more than %d animals and he lives in %s\n", name, cantAnimals, city)
	fmt.Printf("%s has more than %d animals and he lives in %s\n", name, cantAnimals, city)
	// Sprinttf
	message := fmt.Sprintf("%s has more than %d animals and he lives in %s\n", name, cantAnimals, city)
	fmt.Println(message)
	//DataType
	fmt.Printf("Hello message: %T\n", message)
	fmt.Printf("Cant animals: %T\n", cantAnimals)
}
