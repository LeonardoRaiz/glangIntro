package Exercicios

import (
	"fmt"
	"math"
)

func Exercicio7() {
	var raio float64

	fmt.Println("Informe o Raio: ")
	fmt.Scanf("%f", &raio)
	fmt.Scanln()

	area := math.Pi * math.Pow(raio, 2)

	fmt.Printf("A área do círculo é de: %.2f\n", area)
}
