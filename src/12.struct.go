package main

import "fmt"

type car struct {
	brand string
	year  int
}

func structFunction() {
	myCar := car{brand: "Tesla", year: 2022}
	fmt.Println(myCar)

	//Other way to do
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
