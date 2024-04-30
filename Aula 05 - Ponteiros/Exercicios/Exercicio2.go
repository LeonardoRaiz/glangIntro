package Exercicios

/*
Exercício 2: Troca de Valores via Ponteiros
Escreva uma função em Go que aceite dois ponteiros para números inteiros e troque os valores desses dois números. Use esta função em main para trocar os valores de duas variáveis e imprima os resultados antes e depois da troca.
*/

func Exercicio2(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
