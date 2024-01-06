package Exercicios

import "fmt"

func Exercicio5() {
	var salario, perc float64

	fmt.Println("Informe o salário: ")
	fmt.Scanf("%f", &salario)
	fmt.Scanln()
	fmt.Println("Informe o percentual: ")
	fmt.Scanf("%f", &perc)
	fmt.Scanln()

	perc = perc / 100

	salario += salario * perc

	fmt.Println("O novo salário do funcionário é de: ", salario)
}
