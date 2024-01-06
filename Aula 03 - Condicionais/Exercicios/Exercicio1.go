package Exercicios

import "fmt"

func Exercicio1() {
	var numero1 int
	fmt.Print("Informe um número para verificar paridade: ")
	fmt.Scanln(&numero1)
	if numero1%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é ímpar.")
	}
}
