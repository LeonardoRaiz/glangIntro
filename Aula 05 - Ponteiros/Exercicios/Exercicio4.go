package Exercicios

/*
Exercício 4: Ponteiros e Funções
Defina uma função que receba um ponteiro para uma variável inteira e aumente o valor da variável apontada em 10. Utilize essa função no método main para modificar uma variável e exiba seu valor antes e após a chamada da função.
*/

func Exercicio4(x *int) {
	*x += 10
}
