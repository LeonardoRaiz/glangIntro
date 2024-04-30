package Exercicios

import "fmt"

/*
Exercício 1: Manipulação de Maps
Crie um map para armazenar as informações de um livro, incluindo título, autor e ano de publicação. Adicione três livros a este map e implemente:

Uma função para adicionar um novo livro.
Uma função para deletar um livro pelo título.
Uma função para imprimir todos os livros armazenados no map.
*/

func Exercicio1() {
	livros := make(map[string]map[string]string)

	addLivro := func(titulo, autor, ano string) {
		livros[titulo] = map[string]string{"autor": autor, "ano": ano}
	}

	deleteLivro := func(titulo string) {
		delete(livros, titulo)
	}

	printLivros := func() {
		for titulo, info := range livros {
			fmt.Println("Título:", titulo, ", Autor:", info["autor"], ", Ano:", info["ano"])
		}
	}

	// Adicionando livros
	addLivro("Go Programming", "John Doe", "2012")
	addLivro("Advanced Go", "Jane Smith", "2015")
	addLivro("Learning Go", "Emily Davis", "2020")

	// Imprimindo livros
	printLivros()

	// Deletando um livro
	deleteLivro("Advanced Go")

	// Imprimindo livros após deleção
	printLivros()
}
