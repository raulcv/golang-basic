package main

import (
	"fmt"
)

type laptop struct {
	ram   int
	disk  int
	brand string
}

func (p laptop) ping() {
	fmt.Println(p.brand, "Pong")
}
func (p *laptop) duplicateRam() {
	p.ram = p.ram * 2
}

func structPunteroFunction() {
	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(b)
	fmt.Println(*b)

	mypc := laptop{ram: 16, disk: 120, brand: "MACBOOK"}
	mypc.ping()

	fmt.Println(mypc)
	mypc.duplicateRam()
	fmt.Println(mypc)
	mypc.duplicateRam()
	fmt.Println(mypc)

}
