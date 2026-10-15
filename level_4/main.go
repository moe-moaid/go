package main

import ("fmt")

// We wanna allow users to add students, list students, calculate a student, and look up a student

type Student struct {
	name string
	age int
	grade float64
}

var allStudents []Student;


func addStudent(student Student) {
	fmt.Println(student)
	append(allStudents, student)
}

func main() {
	addStudent(Student {name: "moe", age: 25, grade: 66.4})
	fmt.Println(allStudents)
}
