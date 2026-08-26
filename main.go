package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}


func findLargestNumber(arr []int)([]int) {
	var slice = arr[2 : 8]
	return slice
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
	var s = []int{1,3,2,4,5,6,6,7}
	fmt.Println(findLargestNumber(s), len(s), cap(s))
	s = s[2:6]
	fmt.Println("New slice dimentions", len(s), cap(s))
	workWithStructs()
}
