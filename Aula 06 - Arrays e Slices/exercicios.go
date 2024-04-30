package main

import (
	"fmt"
	"moduloArray/Exercicios"
)

func main() {

	fmt.Println("")
	fmt.Println("--- Exercício 1 ---")
	fmt.Println("")
	Exercicios.Exercicio1()

	fmt.Println("")
	fmt.Println("--- Exercício 2 ---")
	fmt.Println("")
	Exercicios.Exercicio2()

	fmt.Println("")
	fmt.Println("--- Exercício 3 ---")
	fmt.Println("")
	Exercicios.Exercicio3()

	fmt.Println("")
	fmt.Println("--- Exercício 4 ---")
	fmt.Println("")
	Exercicios.Exercicio4()

	fmt.Println("")
	fmt.Println("--- Exercício 5 ---")
	fmt.Println("")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	Exercicios.Exercicio5(slice)
	fmt.Println(slice)
}
