package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
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

func multiplicationTable() {
	var n int
	limit := 10
	fmt.Println("Enter a number to multiply")
	fmt.Scan(&n)
	for i :=1; i <= limit; i++ {
		fmt.Printf("%d x %d = %d \n", n, i, n*i)
	}
}

func countToN() {
	var end int
	fmt.Println("Enter a number to count to")
	fmt.Scan(&end)
	for i := 1; i <= end; i++ {
		fmt.Printf("%d, ", i)
	}	
}

func sumOf1ToN () {
	var input int
	sum := 0
	fmt.Println("Enter a number to sum to")
	fmt.Scan(&input)
	for i := 1; i <= input; i++ {
		sum += i
	}
	fmt.Println("the sum is: ", sum)
}

func fizzbuzz () {
	limit := 100
	for i := 0; i<= limit; i++ {
		if i % 5 == 0 && i % 3 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func guessTheNumber() {	
	rn := rand.Intn(1000)
	var input int
	fmt.Println("Guess the number")
	fmt.Scan(&input)
	for input != rn {
		if input > rn {
			fmt.Println("Lower")
			fmt.Scan(&input)
		} else {
			fmt.Println("Higher")
			fmt.Scan(&input)
		}
	}
	fmt.Println("Correct!")
}

func main() {
	/*	printHelloWorld()
	evenOdd()
	calculator()
	negativeVsPositive()
	ageCalc()
	multiplicationTable()
	countToN()
	sumOf1ToN()
	fizzbuzz()
	guessTheNumber()
	*/
	guessTheNumber()	
}
