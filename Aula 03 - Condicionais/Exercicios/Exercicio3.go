package Exercicios

import "fmt"

func Exercicio3() {
	var idade int
	fmt.Print("Informe a idade: ")
	fmt.Scanln(&idade)
	switch {
	case idade >= 60:
		fmt.Println("Idoso.")
	case idade >= 18:
		fmt.Println("Adulto.")
	case idade >= 13:
		fmt.Println("Adolescente.")
	default:
		fmt.Println("Criança.")
	}
}
