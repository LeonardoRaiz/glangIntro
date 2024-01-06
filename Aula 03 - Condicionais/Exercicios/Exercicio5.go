package Exercicios

import "fmt"

func Exercicio5() {
	var lado1, lado2, lado3 int
	fmt.Print("Informe os três lados do triângulo separados por espaços: ")
	fmt.Scanf("%d %d %d", &lado1, &lado2, &lado3)
	fmt.Scanln()
	switch {
	case lado1 == lado2 && lado2 == lado3:
		fmt.Println("Triângulo equilátero.")
	case lado1 == lado2 || lado2 == lado3 || lado1 == lado3:
		fmt.Println("Triângulo isósceles.")
	default:
		fmt.Println("Triângulo escaleno.")
	}
}
