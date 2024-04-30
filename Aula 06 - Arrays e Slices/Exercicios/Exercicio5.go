package Exercicios

/*
Exercício 5: Funções com Slices
Escreva uma função em Go chamada modifySlice que aceita um slice de inteiros e modifica cada elemento do slice, adicionando a ele seu próprio índice (por exemplo, ao elemento no índice 0 será adicionado 0, ao elemento no índice 1 será adicionado 1, e assim por diante). Use esta função em main para modificar um slice, imprima o slice antes e depois da chamada da função.
*/

func Exercicio5(slice []int) {
	for i := range slice {
		slice[i] += i
	}
}
