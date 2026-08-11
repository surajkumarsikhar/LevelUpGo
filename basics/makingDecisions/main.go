package main

import "fmt"

func getGrade(score int) string {
	grade := ""

	// grade already exists, so assign with = (grade = "A"), not :=
	// Check score with if/else if/else and set grade to "A"-"D" or "F"

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	return grade
}

func main() {
	fmt.Println(getGrade(95), getGrade(85), getGrade(72), getGrade(65), getGrade(50))
}