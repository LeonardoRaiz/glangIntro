package Exercicios

import "fmt"

/*
Exercício 1: Inicialização e Acesso
Crie um programa em Go que inicialize um array de inteiros de tamanho 10 com valores de 0 a 9. Após a inicialização, imprima cada valor do array multiplicado por 2.
*/

func Exercicio1() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i] * 2)
	}
	//ou
	for _, value := range array {
		fmt.Println(value * 2)
	}
}
