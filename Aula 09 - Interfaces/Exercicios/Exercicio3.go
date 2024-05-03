package Exercicios

import (
	"fmt"
	"math"
)

/*
Exercício 3: Expansão de Interface
Adicione um novo método à interface forma chamado perimetro() que retorna um float64. Implemente esse método para retangulo e circulo. Para o círculo, o perímetro pode ser calculado como
2π×raio e para o retângulo como
2×(altura+largura). Teste estas implementações usando uma função que escreve o perímetro da forma na tela.
*/

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma1 interface {
	area() float64
	perimetro() float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (r retangulo) perimetro() float64 {
	return 2 * (r.altura + r.largura)
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func Exercicio3() {
	r := retangulo{10, 15}
	c := circulo{10}

	fmt.Printf("Perímetro do retângulo: %0.2f\n", r.perimetro())
	fmt.Printf("Perímetro do círculo: %0.2f\n", c.perimetro())
}
