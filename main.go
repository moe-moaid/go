package main

import (
	"fmt"
	"time"
	"strconv"
)

func printHelloWorld() {
	fmt.Print("Hello, World!\n")
}

func evenOdd() {
	var input int
	fmt.Print("enter a number")
	fmt.Scan(&input)
	var result int = input % 2
	if(result == 0){
		fmt.Print("Even")
	} else {
		fmt.Print("Odd!")
	}
}


func calculator() {
	var first_input, second_input int
	var operation string
	fmt.Print("enter first number")
	fmt.Scan(&first_input)
	fmt.Print("enter second number")
	fmt.Scan(&second_input)
	fmt.Print("enter choose the operation you want to perform")
	fmt.Scan(&operation)
	if(operation == "+"){
		fmt.Println(first_input + second_input)
	} else if (operation == "-") {
		fmt.Println(first_input - second_input)
	} else if (operation == "/") {
		fmt.Println(first_input / second_input)
	} else if (operation == "x") {
		fmt.Println(first_input * second_input)
	}
}

func negativeVsPositive() {
	var input int
	fmt.Print("Enter a number to test for its sign")
	fmt.Scan(&input)
	if(input < 0){
		fmt.Printf("%d is a negative number", input)
	} else if (input > 0) {
		fmt.Printf("%d is a positive number", input)
	} else {
		fmt.Print("the number is zero")
	}
}

func ageCalc() {
	var input int
	current_year, _ := strconv.Atoi(time.Now().Format("2006"))
	fmt.Println("Enter your birth year")
	fmt.Scan(&input)
	fmt.Printf("your age is %d", current_year - input)
}

func main() {
	/*	printHelloWorld()
	evenOdd()
	calculator()
	negativeVsPositive()
	*/
	ageCalc()
}
