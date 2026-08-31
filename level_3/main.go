package main

import ("fmt")

func findLargestNumber(input []int)(largest int) {
	largest = 0
	for i:= 0; i < len(input); i++ {
		if input[i] > largest {
			largest = input[i]
		} 
	}
	return largest
}

func findAverage(input []int)(int){
	sum := 0
	for i := 0; i< len(input); i++ {
		sum += input[i]
	}
	return sum
}

func evenNumbersCount (input []int)(int) {
	var count int
	for i := 0; i < len(input); i++ {
		if input[i] %  2 == 0 {
			count += 1
		}
	}
	return count
}

func main(){
	numbers := []int {1,2,3,4,5,6,7,8,9,10}
	var large int = findLargestNumber(numbers)
	fmt.Println(large)
	var average int = findAverage(numbers)
	fmt.Println(average)
	var count int = evenNumbersCount(numbers)
	fmt.Println(count)
}
