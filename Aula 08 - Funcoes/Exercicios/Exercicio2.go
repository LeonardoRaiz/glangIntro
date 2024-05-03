package Exercicios

import "fmt"

/*
Exercício 2: Função Variádica para Média
Crie uma função variádica chamada media que recebe um número indefinido de parâmetros do tipo float64 e retorna a média desses números. Garanta que a função trate o caso onde nenhum número é passado, retornando 0.
*/
func media(numeros ...float64) float64 {
	total := 0.0
	if len(numeros) == 0 {
		return 0
	}
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func Exercicio2() {
	fmt.Println("Média:", media(7.0, 9.0, 5.5))
}
