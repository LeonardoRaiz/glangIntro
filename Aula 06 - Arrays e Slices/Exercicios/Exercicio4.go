package Exercicios

/*
Exercício 4: Explorando Capacidade de Slices
Crie um slice usando make com um tipo de dados string, uma capacidade inicial de 5 e tamanho 0. Use um loop para adicionar cinco strings ao slice. Após cada adição, imprima a capacidade e o comprimento do slice. Continue adicionando mais cinco strings e observe como a capacidade do slice muda. Discuta ou comente no código sobre como Go gerencia a capacidade de slices dinâmicos.
*/
import "fmt"

func Exercicio4() {
	slice := make([]string, 0, 5)

	for i := 0; i < 5; i++ {
		slice = append(slice, "A")
		fmt.Println("Comprimento = ", len(slice))
		fmt.Println("Capacidade = ", cap(slice))
	}

	for i := 0; i < 5; i++ {
		slice = append(slice, "A")
		fmt.Println("Comprimento = ", len(slice))
		fmt.Println("Capacidade = ", cap(slice))
	}
}
