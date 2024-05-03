package Exercicios

/*
Exercício 1: Expansão da Função de Cálculos Matemáticos
Modifique a função calculosMatematicos para que ela retorne também o produto e a divisão (como um float64) dos dois números passados como argumento, além da soma e subtração já implementadas.
*/

import "fmt"

func calculosMatematicos(n1, n2 int) (soma, subtracao, produto int, divisao float64) {
	soma = n1 + n2
	subtracao = n1 - n2
	produto = n1 * n2
	divisao = float64(n1) / float64(n2)
	return
}

func Exercicio1() {
	soma, subtracao, produto, divisao := calculosMatematicos(10, 2)
	fmt.Println("Soma:", soma, "Subtração:", subtracao, "Produto:", produto, "Divisão:", divisao)
}
