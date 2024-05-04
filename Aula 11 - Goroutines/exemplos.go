package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Concorrência != Paraleleismo
	go escrever("Olá mundo") //goroutine
	escrever("Programando em Go!")

	//Sincronizar goroutine com waitgroup
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Texto no Waitgroup 1")
		waitGroup.Done()
	}()

	go func() {
		escrever("Texto no Waitgroup 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	// Mas tem um jeito melhor... Channels
	canal := make(chan string)
	go escreverChannel("Olá Meu canal", canal)
	fmt.Println("Depois da função escreverChannel começar ser executada")
	for {
		msg, aberto := <-canal // esperando que o canal receba o valor
		if !aberto {
			break
		}
		fmt.Println(msg)
	} // pode acontecer um deadlock

	//refatorar o for acima
	for message := range canal {
		fmt.Println(message)
	}
}

func escrever(texto string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

// nova função com canais
func escreverChannel(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
