package Exercicios

import "fmt"

func Exercicio3() {
	var nota, nota1, nota2 float64

	fmt.Println("Informe a primeira nota: ")
	fmt.Scanf("%f", &nota)
	fmt.Scanln()
	fmt.Println("Informe a segunda nota: ")
	fmt.Scanf("%f", &nota1)
	fmt.Scanln()
	fmt.Println("Informe a terceira nota: ")
	fmt.Scanf("%f", &nota2)
	fmt.Scanln()

	media := (nota*5 + nota1*3 + nota2*2) / (5 + 3 + 2)
	fmt.Println("A média ponderada é igual a: ", media)
}
