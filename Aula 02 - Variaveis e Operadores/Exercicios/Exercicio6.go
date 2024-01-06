package Exercicios

import "fmt"

func Exercicio6() {
	var base, altura float64

	fmt.Println("Informe a base do triângulo: ")
	fmt.Scanf("%f", &base)
	fmt.Scanln()

	fmt.Println("Informe a altura do triângulo: ")
	fmt.Scanf("%f", &altura)
	fmt.Scanln()

	var area float64 = (base * altura) / 2

	fmt.Println("A área do triângulo é de: ", area)
}
