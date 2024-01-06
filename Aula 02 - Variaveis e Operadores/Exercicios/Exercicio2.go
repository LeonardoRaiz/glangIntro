package Exercicios

import "fmt"

func Exercicio2() {
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

	media := (nota + nota1 + nota2) / 3

	fmt.Println("A média aritimética é igual a: ", media)

}
