package Exercicios

import "fmt"

func Exercicio4() {
	var salario float64

	fmt.Println("Informe o salário do funcionário: ")
	fmt.Scanf("%f", &salario)
	fmt.Scanln()

	salario *= 1.25
	// ou salario += salario * 0.25

	fmt.Println("O novo salário do funcionário é de: ", salario)
}
