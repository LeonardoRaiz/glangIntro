package main

import "fmt"

func main() {
	num := 15
	num2 := 30

	soma := num + num2
	fmt.Println("---- Operadores Aritméticos ----")
	fmt.Printf("Soma entre %d e %d é igual: %d\n", num, num2, soma)
	subtracao := num - num2
	fmt.Printf("Subtração entre %d e %d é igual: %d\n", num, num2, subtracao)
	multiplicacao := num * num2
	fmt.Printf("Multiplicação entre %d e %d é igual: %d\n", num, num2, multiplicacao)
	divisao := num / num2 //erro
	fmt.Printf("Divisão entre %d e %d é igual: %d\n", num, num2, divisao)
	var num3 float64 = 15
	var num4 float64 = 30
	var divisao1 float64 = num3 / num4 //erro
	fmt.Printf("Divisão entre %d e %d é igual: %f\n", num, num2, divisao1)
	resto := num % num2
	fmt.Printf("Resto da divisão entre %d e %d é igual: %d\n", num, num2, resto)

	fmt.Println("---- Operadores de Comparação ----")

	igual := num == num2
	fmt.Printf("Os números são iguais? %t\n", igual)
	diferente := num != num2
	fmt.Printf("Os números são diferentes? %t\n", diferente)
	maior := num > num2
	fmt.Printf("O número %d é maior que %d? %t\n", num, num2, maior)
	menor := num < num2
	fmt.Printf("O número %d é menor que %d? %t\n", num, num2, menor)
	maiorOuIgual := num >= num2
	fmt.Printf("O número %d é maior ou igual que %d? %t\n", num, num2, maiorOuIgual)
	menorOuIgual := num <= num2
	fmt.Printf("O número %d é menor ou igual que %d? %t\n", num, num2, menorOuIgual)

	fmt.Println("---- Operadores Lógicos ----")

	condicao1 := true
	condicao2 := false
	resultadoE := condicao1 && condicao2
	fmt.Printf("O resultado do operador E = %t\n", resultadoE)
	resultadoOu := condicao1 || condicao2
	fmt.Printf("O resultado do operador OU = %t\n", resultadoOu)
	negacao := !condicao1
	fmt.Printf("O resultado do operador NOT = %t\n", negacao)

	fmt.Println("---- Operadores Lógicos ----")

	num += 5 // Equivalente a: num = num + 5
	fmt.Println("Novo valor de num: ", num)
	num -= 5 // Equivalente a: num = num - 5
	fmt.Println("Novo valor de num: ", num)
	num *= 2 // Equivalente a: num = num * 2
	fmt.Println("Novo valor de num: ", num)
	num /= 5 // Equivalente a: num = num / 5
	fmt.Println("Novo valor de num: ", num)

}
