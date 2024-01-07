package Exercicios

import "fmt"

func Exercicio1() {
	var numero int
	fmt.Print("Informe um número para a tabuada: ")
	fmt.Scanln(&numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
