package main

import ("fmt")

// a func to find the lartest number in a slice
func findLargestNumber(input []int)(largest int) {
	largest = 0
	for i:= 0; i < len(input); i++ {
		if input[i] > largest {
			largest = input[i]
		} 
	}
	return largest
}

// a func to find the average of an slice
func findAverage(input []int)(int){
	sum := 0
	for i := 0; i< len(input); i++ {
		sum += input[i]
	}
	return sum
}

// a func to find the number of even numbers in a slice
func evenNumbersCount (input []int)(int) {
	var count int
	for i := 0; i < len(input); i++ {
		if input[i] %  2 == 0 {
			count += 1
		}
	}
	return count
}

// a func to reverce a slice
func reverceSlice (input []int) ([]int) {
	var newArr [] int
	for i := len(input) -1 ; i > 0; i-- {
		newArr = append(newArr, input[i])
	}
	return newArr
}

func main(){
	numbers := []int {1,2,3,4,5,6,7,8,9,10} // the main slice that willbe used in all funcs below
	var large int = findLargestNumber(numbers)
	fmt.Println(large)
	var average int = findAverage(numbers)
	fmt.Println(average)
	var count int = evenNumbersCount(numbers)
	fmt.Println(count)
	var invert []int = reverceSlice(numbers)
	fmt.Println(invert)
}
