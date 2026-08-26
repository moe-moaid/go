package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}


func findLargestNumber(arr []int)(int) {
	var result int
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > result {
			result = arr[i]
		} else if arr[i+1] > result {
			result = arr[i+1]
		}
	}
	return result
}

func workWithStructs() {
	pp := Person{"Moe", 47}
	point := &pp
	fmt.Println(pp.name, pp.age)
	point.name = "jake"
	fmt.Println(pp.name, pp.age)
	var something int = 5
	var pointer = &something
	fmt.Println(*pointer)
	*pointer = 88
	fmt.Println(*pointer)
}

func main() {
	var s = []int{1,3,2,12,5,23,6,7}
	fmt.Printf("largest number is:: %d \n", findLargestNumber(s))
	workWithStructs()
}
