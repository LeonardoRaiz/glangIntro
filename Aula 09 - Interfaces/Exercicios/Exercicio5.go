package Exercicios

import "fmt"

/*
Exercício 5: Polimorfismo com Slices de Interfaces
Crie um slice do tipo forma e adicione um retangulo, um circulo e um triangulo a ele. Utilize um loop para iterar sobre este slice e chame a função escreverArea() para cada elemento do slice, demonstrando polimorfismo.
*/

type forma2 interface {
	area() float64
}

func Exercicio5() {
	formas := []forma{
		retangulo{10, 15},
		circulo{10},
		triangulo{5, 7},
	}

	for _, forma := range formas {
		escreverArea1(forma)
	}
}

func escreverArea1(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}
