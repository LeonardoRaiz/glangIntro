package Exercicios

import "fmt"

/*

Exercício 1: Manipulação Básica de Ponteiros
Crie um programa em Go onde duas variáveis inteiras a e b são declaradas. Defina a como 5 e b como o endereço de a. Modifique o valor de a usando b. Imprima o novo valor de a.

*/

func Exercicio1() {
	var a int = 5
	var b *int = &a
	fmt.Println(a)
	*b = 10
	fmt.Println(a)
}
