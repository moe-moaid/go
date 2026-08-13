package main

import "fmt"

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
	fmt.Print("enter first number")
	fmt.Scan(&second_input)
	fmt.Print("enter choose the operation you want to perform")
	fmt.Scan(&operation)
	if(operation == "+"){
		fmt.Printf(first_input + second_input)
	} else if (operation == "-") {
		fmt.Printf(first_input - second_input)
	} else if (operation == "/") {
		fmt.Printf(first_input / second_input)
	} else if (operation == "x") {
		fmt.Printf(first_input * second_input)
	}
}
	
func main() {
	printHelloWorld()
	evenOdd()
	calculator()
}
