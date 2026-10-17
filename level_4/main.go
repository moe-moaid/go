package main

import ("fmt")

// We wanna allow users to add students, list students, calculate average grades, and look up a student

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

func calcAvg() float64 {
	var sum float64
	for i := 0; i < len(allStudents); i++ {
		sum += allStudents[i].grade
	}
	return sum/float64(len(allStudents))
}

func main() {
	testStudent := Student  {name: "moe", age: 25, grade: 66.4}
	addStudent(testStudent)
	addStudent(testStudent)
	addStudent(testStudent)
	addStudent(testStudent)
	fmt.Println("allStudents slice:: ", allStudents)
	listStudents()
	fmt.Println("average grades:: ", calcAvg())
}
