package Exercicios

import "fmt"

/*
Exercício 3: Endereços e Ponteiros
Elabore um programa que declara uma variável inteira x e um ponteiro p para inteiro. Atribua o endereço de x a p e imprima o endereço de x e o valor de p. Modifique o valor de x através de p e imprima o novo valor de x.
*/

func Exercicio3() {
	var x int = 10
	var p *int = &x
	fmt.Println("Endereço de X = ", &x, "\nValor de P = ", *p)
	*p = 50
	fmt.Println("Novo valor de X = ", x)
}
