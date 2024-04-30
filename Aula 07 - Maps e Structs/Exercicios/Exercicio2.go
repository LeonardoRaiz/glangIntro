package Exercicios

import "fmt"

/*
Exercício 2: Structs Aninhados
Defina uma struct para representar um Carro, que inclui campos para marca, modelo, e uma struct aninhada Especificacoes, que deve ter cilindradas e potência do motor. Crie e inicialize uma instância desta struct, e imprima todas as suas informações.
*/

type Especificacoes struct {
	Cilindradas int
	Potencia    int
}

type Carro struct {
	Marca         string
	Modelo        string
	Especificacao Especificacoes
}

func Exercicio2() {
	meucarro := Carro{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Especificacao: Especificacoes{
			Cilindradas: 1800,
			Potencia:    140,
		},
	}
	fmt.Printf("Carro: %s %s\n", meucarro.Marca, meucarro.Modelo)
	fmt.Printf("Especificações: %d cc, %d HP\n", meucarro.Especificacao.Cilindradas, meucarro.Especificacao.Potencia)
}
