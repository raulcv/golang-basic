package main

import (
	"fmt"

	"github.com/raulcv/go-basic/src/myownpkg"
	pkrace "github.com/raulcv/go-basic/src/race"
)

func privatePublicAccessFuncs() {

	var myPublicCar pkrace.CarPublic
	myPublicCar.Brand = "Mazda"
	myPublicCar.Year = 2029
	fmt.Println(myPublicCar)
	myownpkg.PrintMessage("HEY! WHATS UPP")
}
