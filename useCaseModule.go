package main

import "fmt"

//Head é o primeiro elemento, Tail é o ultimo elemento
type GradesList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	//Value guarda um objeto na memoria do tipo grade
	Value Grades
	//Next é o próximo elemento e é um ponteiro do Node
	Next *Node
}

//Append adiciona um novo valor na GradesList
func (gl *GradesList) Append(grades Grades) {
	//node recebe um novo Value to tipo Grades
	node := &Node{Value: grades}
	//Verifica se a gradesList está vazia e aponto o Head dela para esse Value do Node
	if gl.Head == nil {
		gl.Head = node
	}
	// Verifica se o Tail está vazio e se estiver eu pego o Next dele e aponto para o novo Value do Node
	if gl.Tail != nil {
		gl.Tail.Next = node
	}
	//Tail passa a ser o novo node
	gl.Tail = node

}

//Search faz uma busca do nome do estudante para achar sua Grade
func (gl *GradesList) Search(studentName string) Grades {
	//Começa com o primeiro elemento da GradesList
	node := gl.Head
	//Quando chegar no break o looping para
	for node != nil {
		if node.Value.StudentName == studentName {
			break
		}
		//próximo elemento
		node = node.Next
	}

	if node != nil {
		return node.Value
	}
	//Não encontramos ninguem com esse studentName
	return Grades{}

}

func (gl *GradesList) Remove(studentName string) {
	//Validar se o primeiro item da lista já e o que eu to buscando
	if gl.Head.Value.StudentName == studentName {
		gl.Head = gl.Head.Next
		return
	}

	//busca na lista
	previousValue := gl.Head
	node := gl.Head.Next

	//Iterar entre os nodes ate encontrar para remover
	for node != nil {
		//encontra meu nome
		if node.Value.StudentName == studentName {
			// previousValue anterior vai apontar para o  proximo
			previousValue.Next = node.Next
			break
		}
		//troca
		previousValue = node
		node = node.Next
	}

	//validar se o node que eu cenontrei é o tail e apontar ele para o anterior
	if gl.Tail == node {
		gl.Tail = previousValue
	}

}

//Quero iterar entre a Gradeslist e mostrar os resutados encontrados
func (gl *GradesList) Display() {
	node := gl.Head
	//Cada um dos nodes tem o próximo valor e o ultimo é o nil então ele sai do looping e acaba
	for node != nil {
		fmt.Println(node.Value.Grade)
		node = node.Next
	}

}
