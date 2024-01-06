package Exercicios

import "fmt"

func Exercicio1() {
	var num, num1, num2, num3 int
	/*
		fmt.Println("Informe os valores( O espaço separa os valores ): ")
		fmt.Scanf("%d %d %d %d", &num, &num1, &num2, &num3)
		soma := num + num1 + num2 + num3
		fmt.Printf("A soma entre os valores %d + %d + %d + %d é igual a: %d\n", num, num1, num2, num3, soma)
	*/

	fmt.Printf("Informe o valor: ")
	fmt.Scanf("%d", &num)
	fmt.Scanln() // Limpa o buffer

	fmt.Printf("Informe o valor: ")
	fmt.Scanf("%d", &num1)
	fmt.Scanln() // Limpa o buffer

	fmt.Printf("Informe o valor: ")
	fmt.Scanf("%d", &num2)
	fmt.Scanln() // Limpa o buffer

	fmt.Printf("Informe o valor: ")
	fmt.Scanf("%d", &num3)
	fmt.Scanln()

	soma := num + num1 + num2 + num3
	fmt.Printf("A soma entre os valores %d + %d + %d + %d é igual a: %d\n", num, num1, num2, num3, soma)

}
