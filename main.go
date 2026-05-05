package main

import (
	"fmt"
	"learning-struct/students"
)


func main() {

	s1 := students.Student{
		Name:  "Alabi Sunday",
		Age:   23,
		Score: []float64{90.0, 100, 99.9},
	}
	s2 := students.Student{
		Name:  "Adaobi uche",
		Age:   34,
		Score: []float64{40.0, 20.4, 90.1},
	}
	s3 := students.Student{
		Name:  "Emeka Adekunle",
		Age:   34,
		Score: []float64{40.0, 20.4, 90.1},
	}
	//Create new type of students with their corresponding score
	studentList := []students.Student{s1, s2, s3}
	fmt.Printf("%-20s | %-15s | %-10s\n", "Name", "Average Score", "Status")
	fmt.Println("------------------------------------------------------------")

	for i := 0; i < len(studentList); i++ {
		s := studentList[i]
		average := s.AverageScore()
		status := "fail"

		if s.HasPassed() {
			status = "passed"
		}
		fmt.Printf("%-20s | %-15.2f | %-10s\n", s.Name, average, status)
	}

}
