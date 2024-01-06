package Exercicios

import (
	"fmt"
	"math"
)

func Exercicio8() {
	var numero float64

	fmt.Println("Informe o número: ")
	fmt.Scanf("%f", &numero)
	fmt.Scanln()

	quad := math.Pow(numero, 2)
	cubo := math.Pow(numero, 3)
	raizQ := math.Sqrt(numero)
	raizC := math.Pow(numero, (1.0 / 3.0))

	fmt.Printf("O número digitado ao quadrado: %f\n"+
		"O número digitado ao cubo: %f\n"+
		"A raiz quadrada do número digitado: %f\n"+
		"A raiz cúbica de número digitado: %f\n", quad, cubo, raizQ, raizC)
}
