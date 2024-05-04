package Exercicios

import (
	"fmt"
	"time"
)

/*
Exercício 1: Goroutines Simples
Crie um programa que inicia três goroutines separadas, cada uma chamando a função escrever com mensagens diferentes ("Go Routine 1", "Go Routine 2", "Go Routine 3"). Cada goroutine deve executar a função escrever apenas duas vezes. Observe e discuta a ordem das saídas.
*/

func Exercicio1() {
	go escrever("Go Routine 1")
	go escrever("Go Routine 2")
	go escrever("Go Routine 3")
	time.Sleep(time.Second)
}

func escrever(texto string) {
	for i := 0; i < 2; i++ {
		fmt.Println(texto)
	}

}
