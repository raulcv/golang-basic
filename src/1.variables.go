package main

import "fmt"

func variables() {
	//Constants
	const piNumber float64 = 3.14
	const piNumberTwo = 3.1416

	fmt.Println("PI VALUE: ", piNumber)
	fmt.Println("PI VALUE 2: ", piNumberTwo)
	//INT VALS
	base := 13
	var height int = 14
	var area int = 12

	println("Area: ", base)
	println("Heiht: ", height)
	println("Area: ", area)
	//ZERO VALUES
	var a int
	var b float64
	var c string
	var d bool
	println(a, b, c, d)
	//Area
	const baseCuadrado = 12
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)
}
