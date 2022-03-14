package main

import (
	"fmt"
)

func makeMapFunction() {
	m := make(map[string]int)
	m["Raul"] = 14
	m["Señor de Sipan"] = 15
	fmt.Println(m)
	//Record map value
	for i, v := range m {
		fmt.Println(i, v)
	}
	//Find a value into the map
	valueFound, status := m["Raul"]
	// fmt.Println(valueFound, i)
	if status == false {
		fmt.Println("value not Found")
	} else {
		fmt.Println(valueFound)
	}
}
