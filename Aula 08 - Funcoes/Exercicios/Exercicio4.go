package Exercicios

import "fmt"

/*
Exercício 4: Funções Anônimas em Loop
Escreva um programa que utilize uma função anônima em um loop for para imprimir os números de 1 a 5. Utilize a função anônima dentro do loop para realizar a impressão.
*/

func Exercicio4() {
	for i := 1; i <= 5; i++ {
		func(n int) {
			fmt.Println(n)
		}(i)
	}
}
