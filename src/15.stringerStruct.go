package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (p *pc) Stringer() string {
	return fmt.Sprintf("I have %d GB of Ram, %d GB of Disk and it's a %s", p.ram, p.disk, p.brand)
}

func StringerStringFunction() {
	mypc := pc{ram: 16, disk: 120, brand: "MACBOOK"}
	fmt.Println(mypc)
	fmt.Println(mypc.Stringer())
}
