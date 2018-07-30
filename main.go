package main

import (
	"fmt"
)

func main() {
	grades := []int32{73, 67, 38, 33}
	result := gradingStudents(grades)
	for _, grade := range result {
		fmt.Println(grade)
	}
}
