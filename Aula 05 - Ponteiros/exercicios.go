package main

import (
	"fmt"
	"moduloPonteiro/Exercicios"
)

func main() {

	fmt.Println("")
	fmt.Println("--- Exercício 1 ---")
	fmt.Println("")
	Exercicios.Exercicio1()

	fmt.Println("")
	fmt.Println("--- Exercício 2 ---")
	fmt.Println("")
	a := 5
	b := 10
	fmt.Println("Antes da troca a=", a, "b=", b)
	Exercicios.Exercicio2(&a, &b)
	fmt.Println("Depois da troca a=", a, "b=", b)

	fmt.Println("")
	fmt.Println("--- Exercício 3 ---")
	fmt.Println("")
	Exercicios.Exercicio3()

	fmt.Println("")
	fmt.Println("--- Exercício 4 ---")
	fmt.Println("")
	c := 10
	fmt.Println("Valor antes do incremento C = ", c)
	Exercicios.Exercicio4(&c)
	fmt.Println("Valor depois do incremento C = ", c)

	fmt.Println("")
	fmt.Println("--- Exercício 5 ---")
	fmt.Println("")
}
