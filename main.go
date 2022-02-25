package main

import (
	"fmt"
)

type GradeAndSubject struct {
	StudentName string
	Grade       float64
	NameSubject string
	Teacher     string
}

func main() {
	gradeAndSubjectList := &GradeAndSubjectList{}

	roberta := GradeAndSubject{"Roberta", 10.00, "POO", "Carlos"}
	greg := GradeAndSubject{"Greg", 6.00, "POO", "Carlos"}
	nando := GradeAndSubject{"Fernando", 8.00, "POO", "Carlos"}
	vinicius := GradeAndSubject{"Vinicius", 2.00, "POO", "Carlos"}

	gradeAndSubjectList.Append(roberta)
	gradeAndSubjectList.Append(greg)
	gradeAndSubjectList.Append(nando)
	gradeAndSubjectList.Append(vinicius)

	//TODO
	search := gradeAndSubjectList.Search("Roberta")
	fmt.Printf("A nota do aluno %s é %.2f na matéria %v do professor %v.",
		search.StudentName, search.Grade, search.NameSubject, search.Teacher)

	gradeAndSubjectList.Delete("greg")

	//gradeAndSubjectList.Display()
}
