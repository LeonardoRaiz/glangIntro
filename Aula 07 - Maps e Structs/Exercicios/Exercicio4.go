package Exercicios

import "fmt"

/*
Exercício 4: Maps Complexos
Elabore um programa que utiliza um map onde as chaves são nomes de estados do Brasil e os valores são outro map que associa o nome de uma cidade a sua população estimada. Adicione pelo menos três estados com três cidades cada. Implemente funções para:

Adicionar uma nova cidade e sua população a um estado.
Remover uma cidade de um estado.
Imprimir a população de um estado, somando as populações de suas cidades.
*/

func Exercicio4() {
	estados := make(map[string]map[string]int)

	addCidade := func(estado, cidade string, populacao int) {
		if estados[estado] == nil {
			estados[estado] = make(map[string]int)
		}
		estados[estado][cidade] = populacao
	}

	removeCidade := func(estado, cidade string) {
		delete(estados[estado], cidade)
	}

	printPopulacaoEstado := func(estado string) {
		total := 0
		for _, pop := range estados[estado] {
			total += pop
		}
		fmt.Println("População Total do Estado", estado, ":", total)
	}

	// Adicionando cidades
	addCidade("São Paulo", "São Paulo", 12300000)
	addCidade("São Paulo", "Campinas", 1214000)
	addCidade("Rio de Janeiro", "Rio de Janeiro", 6748000)
	addCidade("Rio de Janeiro", "Niterói", 513000)

	// Imprimindo populações
	printPopulacaoEstado("São Paulo")
	printPopulacaoEstado("Rio de Janeiro")

	// Removendo uma cidade
	removeCidade("São Paulo", "Campinas")

	// Imprimindo populações após remoção
	printPopulacaoEstado("São Paulo")
}
