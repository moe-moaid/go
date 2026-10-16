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
	allStudents = append(allStudents, student)
	return allStudents
}

func listStudents() {
	for i:= 0; i < len(allStudents); i ++ {
		fmt.Println(allStudents[i])
	}
}

func main() {
	testStudent := Student  {name: "moe", age: 25, grade: 66.4}
	addStudent(testStudent)
	addStudent(testStudent)
	addStudent(testStudent)
	addStudent(testStudent)
	fmt.Println("allStudents slice:: ", allStudents)
	listStudents()
}
