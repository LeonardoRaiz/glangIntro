package Exercicios

import "fmt"

func Exercicio2() {
	var nota float64
	fmt.Print("Informe a nota do aluno: ")
	fmt.Scanln(&nota)
	if nota >= 7 {
		fmt.Println("Aluno aprovado!")
	} else if nota >= 4 && nota < 7 {
		fmt.Println("Aluno em exame.")
	} else {
		fmt.Println("Aluno reprovado.")
	}
}
