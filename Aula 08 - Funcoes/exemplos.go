package main

import "fmt"

// Retorno Nomeado
func calculosMatematicos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// Função Variáticas
func soma2(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Função Recursiva
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	soma, subtracao := calculosMatematicos(3, 4)
	fmt.Println(soma, subtracao)

	totalSoma := soma2(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)

	posicao := 10
	fmt.Println(fibonacci(posicao))

	// Função Anônimas
	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

}
