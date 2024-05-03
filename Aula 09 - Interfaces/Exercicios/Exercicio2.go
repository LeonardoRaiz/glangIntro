package Exercicios

import (
	"fmt"
	"reflect"
)

/*
Exercício 2: Funções de Forma Genérica
Modifique a função generica() para que ela identifique e imprima o tipo do argumento passado usando o pacote reflect. Teste sua função modificada passando diferentes tipos de dados como string, int, bool e retangulo.
*/

func generica(interf interface{}) {
	fmt.Println("Tipo do dado:", reflect.TypeOf(interf), "Valor:", interf)
}

func Exercicio2() {
	generica("Olá, mundo!")
	generica(42)
	generica(true)
	generica(3.1415)
}
