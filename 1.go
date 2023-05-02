package main

import "fmt"

type Retângulo struct {
	largura float64
	altura  float64
}

func calcularArea(r Retângulo) {
	var resul float64
	resul = r.largura * r.altura
	fmt.Println("A área desse retângulo", resul)
}

func main() {
	r := Retângulo{largura: 5, altura: 8.0}
	calcularArea(r)
}
