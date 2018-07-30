package main

// Grading Students
func gradingStudents(grades []int32) []int32 {
	result := []int32{}
	for _, grade := range grades {
		diffToBeRounded := grade%5 - 5
		canBeRounded := -diffToBeRounded < 3
		if canBeRounded && grade >= 38 {
			result = append(result, grade+(-diffToBeRounded))
		} else {
			result = append(result, grade)
		}
	}
	return result
}
