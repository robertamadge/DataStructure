package main

//GradeAndSubjectList guarda o Head que é o primeiro elemento da lista e Tail, o último elemento
type GradeAndSubjectList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	//Value guarda um objeto na memoria do tipo grade
	Value GradeAndSubject
	//Next é o próximo elemento e é um ponteiro do Node
	Next *Node
}

//Append adiciona um novo valor na GradeAndSubjectList(reference)
func (gl *GradeAndSubjectList) Append(gs GradeAndSubject) {
	//node recebe um novo Value to tipo GradesAndSubject
	node := &Node{Value: gs}
	//Verifica se a GradeAndSubjectList está vazia e aponto o Head dela para esse Value do Node
	if gl.Head == nil {
		gl.Head = node
	}
	// Verifica se o Tail não está vazio e pego o Next dele e aponto para o novo node
	if gl.Tail != nil {
		gl.Tail.Next = node
	}
	//Se não, Tail passa a ser o novo node
	gl.Tail = node
}

//Search faz uma busca com o nome do estudante para achar sua nota e subject
func (gl *GradeAndSubjectList) Search(studentName string) GradeAndSubject {
	//Busca pelo começo da lista
	node := gl.Head
	//Se a lista não estiver vazia, verifica se o nome do estudante é igual ao que eu procuro e break
	for node != nil {
		if node.Value.StudentName == studentName {
			break
		}
		//Se não, verifico o próximo elemento
		node = node.Next
	}
	//Retorno o valor
	if node != nil {
		return node.Value
	}
	//Se não encontramos ninguem com esse studentName
	return GradeAndSubject{}
}

//Delete quero deletar a nota do estudante que informar
func (gl *GradeAndSubjectList) Delete(studentName string) {
	//Validar se o primeiro item da lista já e o que eu to buscando, trocar pelo próximo
	if gl.Head.Value.StudentName == studentName {
		gl.Head = gl.Head.Next
		return
	}
	//se não for o primeiro continuo buscando na lista.  previousValue é o primeiro que achamos mas que não queremos
	previousValue := gl.Head
	//o valor seguinte do previousValue
	node := gl.Head.Next

	//Iterar entre os nodes ate encontrar para remover. Checa se o novo valor != nil
	for node != nil {
		//encontra o nome do estudante
		if node.Value.StudentName == studentName {
			// previousValue anterior vai apontar para o  proximo, deletando ele e apontando para o proximo. Metodo de exclusao
			previousValue.Next = node.Next
			break
		}
		// Se não encontrar faço uma troca
		previousValue = node
		node = node.Next
	}
	//validar se o nome que eu encontrei é o tail e apontar ele para o anterior
	if gl.Tail == node {
		gl.Tail = previousValue
	}
}

// //Quero iterar entre a GradeAndSubjectList e mostrar os resutados encontrados -- Tirar comentario ctrl + 7
// func (gl *GradeAndSubjectList) Display() {
// 	node := gl.Head
// 	//Cada um dos nodes tem o próximo valor e o ultimo é o nil então ele sai do looping e acaba
// 	for node != nil {
// 		fmt.Println(node.Value)
// 		node = node.Next
// 	}
// }
