package main

import "fmt"

func arrSliceFunction() {
	//Array
	var arrayValue [10]int
	arrayValue[0] = 1
	arrayValue[1] = 2
	arrayValue[2] = 4
	fmt.Println(arrayValue)
	arrayLenth := len(arrayValue)
	fmt.Println("Lenght of Array: ", arrayLenth)
	capArray := cap(arrayValue)
	fmt.Println("Maximum capacity for Array: ", capArray)
	//Slice
	sliceArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//Slice Methos
	fmt.Println("Value of position 0 array: ", sliceArray[0])
	fmt.Println("Value of array until to three position: ", sliceArray[:3])
	fmt.Println("Value of array beetween 2 - 5: ", sliceArray[2:5])
	fmt.Println("Value of array from Four to max: ", sliceArray[4:])
	//Append
	sliceArray = append(sliceArray, 11, 12, 13, 14, 15)
	fmt.Println("Add value of slice of array: ", sliceArray)

	//Append new slice
	newSliceArray := []int{16, 17, 18, 19, 20}
	sliceArray = append(sliceArray, newSliceArray...)
	fmt.Println("New slice of array: ", sliceArray)
}
