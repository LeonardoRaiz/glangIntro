package Exercicios

import "fmt"

func Exercicio4() {
	var numero int
	fmt.Print("Informe um número para calcular o fatorial: ")
	fmt.Scanln(&numero)

	resultado := 1
	for i := 1; i <= numero; i++ {
		resultado *= i
	}

	fmt.Printf("O fatorial de %d é: %d\n", numero, resultado)
}
