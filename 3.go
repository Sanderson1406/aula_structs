package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func mediaAluno(m Aluno) {
	m.notas = []float64{8.7, 6.8, 4.8, 9.8, 7.0}
	var soma float64
	for i := 0; 1 < len(m.notas); i++ {
		soma += m.notas[i]
	}
	resul := soma / 5
	fmt.Println(resul)
}
