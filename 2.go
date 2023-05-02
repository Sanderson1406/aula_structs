package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	ano    string
}

func main() {
	l := Livro{titulo: "A Guerra dos Tronos", autor: "Goerge", ano: "21/05/2023"}
	fmt.Println("Título: ", l.titulo)
	fmt.Println("Autor: ", l.autor)
	fmt.Println("Data de lançamento: ", l.ano)
}
