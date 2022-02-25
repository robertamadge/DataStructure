package main

import "fmt"

type Grades struct {
	StudentName  string
	Grade        float64
	Registration int
}

func main() {
	gradesList := GradesList{}

	roberta := Grades{"Roberta", 10.00, 123456}
	carlos := Grades{"Carlos", 5.00, 789101}
	greg := Grades{"Greg", 6.00, 121314}

	gradesList.Append(roberta)
	gradesList.Append(carlos)
	gradesList.Append(greg)

	grade := gradesList.Search("roberta")
	fmt.Println("A nota do aluno ", roberta.StudentName, "e sua nota é: ", grade)

	gradesList.Delete(6.00)

	gradesList.Display()
}
