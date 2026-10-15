package main

import ("fmt")

// We wanna allow users to add students, list students, calculate a student, and look up a student

type Student struct {
	name string
	age int
	grade float64
}

var allStudents []Student;


func addStudent(student Student) []Student {
	fmt.Println("student slice: ", student.name)
	allStudents = append(allStudents, student)
	return allStudents
}

func main() {
	testStudent := Student  {name: "moe", age: 25, grade: 66.4}
	fmt.Println(addStudent(testStudent))
	fmt.Println("allStudents slice:: ", allStudents)
}
