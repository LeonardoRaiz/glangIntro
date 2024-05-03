package Exercicios

import "fmt"

/*
Exercício 5: Função Recursiva de Fatorial
Implemente uma função recursiva que calcule o fatorial de um número. Teste a função com diferentes valores para garantir que está funcionando corretamente.
*/

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func Exercicio5() {
	fmt.Println("Fatorial de 5:", fatorial(5))
	fmt.Println("Fatorial de 7:", fatorial(7))
}
