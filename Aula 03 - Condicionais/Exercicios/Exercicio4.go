package Exercicios

import "fmt"

func Exercicio4() {
	var num1, num2, num3 int
	fmt.Print("Informe três números separados por espaços: ")
	fmt.Scanf("%d %d %d", &num1, &num2, &num3)
	fmt.Scanln()
	maximo := num1
	if num2 > maximo {
		maximo = num2
	}
	if num3 > maximo {
		maximo = num3
	}
	fmt.Printf("O número máximo é: %d\n", maximo)
}
