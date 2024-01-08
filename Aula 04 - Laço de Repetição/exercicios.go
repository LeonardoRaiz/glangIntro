package main

import (
	"fmt"
	"moduloRepeticao/Exercicios"
)

func main() {
	i := 1
	for i != 0 {
		fmt.Println("Qual exercício você deseja visualizar?")
		fmt.Println("O valor 0 sair do programa!")
		fmt.Scanln(&i)

		switch i {
		case 1:
			fmt.Println("--- Exercício 1 ---")
			fmt.Println("")
			Exercicios.Exercicio1()
		case 2:
			fmt.Println("")
			fmt.Println("--- Exercício 2 ---")
			fmt.Println("")
			Exercicios.Exercicio2()
		case 3:
			fmt.Println("")
			fmt.Println("--- Exercício 3 ---")
			fmt.Println("")
			Exercicios.Exercicio3()
		case 4:
			fmt.Println("")
			fmt.Println("--- Exercício 4 ---")
			fmt.Println("")
			Exercicios.Exercicio4()
		}
	}

}
