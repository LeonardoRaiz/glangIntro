package Exercicios

import "fmt"

/*
Exercício 3: Herança Simulada com Structs
Crie uma struct Funcionario que incorpora a struct Pessoa do exemplo, e adicione um campo adicional para o departamento. Defina um slice de Funcionario e inicialize-o com vários funcionários. Implemente uma função que lista todos os funcionários de um determinado departamento.
*/

type Pessoa struct {
	Nome, Sobrenome string
	Idade           int
}

type Funcionario struct {
	Pessoa
	Departamento string
}

func listFuncionarios(departamento string, funcionarios []Funcionario) {
	for _, f := range funcionarios {
		if f.Departamento == departamento {
			fmt.Println(f.Nome, f.Sobrenome, "do departamento de", departamento)
		}
	}
}

func Exercicio3() {
	funcionarios := []Funcionario{
		{Pessoa{"João", "Silva", 30}, "TI"},
		{Pessoa{"Maria", "Oliveira", 25}, "Recursos Humanos"},
		{Pessoa{"Pedro", "Costa", 28}, "TI"},
	}

	listFuncionarios("TI", funcionarios)
}
