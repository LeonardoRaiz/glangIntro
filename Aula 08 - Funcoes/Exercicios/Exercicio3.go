package Exercicios

import "fmt"

/*
Exercício 3: Série de Fibonacci Modificada
Adapte a função fibonacci para que ela retorne um slice contendo todos os números da série de Fibonacci até a posição especificada, em vez de retornar um único número.
*/

func fibonacci(posicao int) []int {
	resultados := make([]int, posicao+1)
	if posicao >= 0 {
		resultados[0] = 0
	}
	if posicao >= 1 {
		resultados[1] = 1
	}
	for i := 2; i <= posicao; i++ {
		resultados[i] = resultados[i-1] + resultados[i-2]
	}
	return resultados[:posicao+1]
}

func Exercicio3() {
	fmt.Println("Série Fibonacci:", fibonacci(10))
}
