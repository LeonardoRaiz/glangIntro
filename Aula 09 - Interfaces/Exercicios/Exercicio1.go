package Exercicios

import "fmt"

/*
Exercício 1: Implementação de Interface
Defina uma nova struct chamada triangulo que possui a base e a altura como campos. Implemente o método area() para essa struct, que deverá calcular a área do triângulo (
base∗altura/2). Utilize a função escreverArea() para exibir a área de um objeto triangulo.
*/

type triangulo struct {
	base   float64
	altura float64
}

func (t triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

func Exercicio1() {
	t := triangulo{base: 5, altura: 7}
	escreverArea(t)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type forma interface {
	area() float64
}
