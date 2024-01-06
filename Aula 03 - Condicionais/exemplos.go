package main

import "fmt"

func main() {
	numero := 10

	if numero > 5 {
		fmt.Println("O número é maior que 5.")
	}

	if numero > 0 {
		fmt.Println("O número é positivo.")
	} else if numero < 0 {
		fmt.Println("O número é negativo.")
	} else {
		fmt.Println("O número é zero.")
	}

	fmt.Println("--- Switch Case ---")

	dia := "terça"

	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia útil.")
	case "sábado", "domingo":
		fmt.Println("Fim de semana.")
	default:
		fmt.Println("Dia inválido.")
	}

}
