package Exercicios

/*
Exercício 2: Slices Dinâmicos
Escreva um programa que inicie com um slice vazio de inteiros e, em seguida, use um loop para adicionar os números de 1 a 15 ao slice usando a função append. Imprima o slice antes e depois da adição dos números, bem como sua capacidade em cada etapa.
*/
import "fmt"

func Exercicio2() {
	slice := []int{}

	for i := 1; i <= 15; i++ {
		slice = append(slice, i)
		fmt.Println("Slice:", slice, "Capacidade:", cap(slice))
	}
}
