package Exercicios

import (
	"fmt"
	"sync"
)

/*
Exercício 4: Canais com Capacidade
Crie um canal com buffer (capacidade) e uma goroutine que envia 10 mensagens para este canal. Utilize uma segunda goroutine para receber e imprimir essas mensagens. Use WaitGroup para sincronizar o término de ambas as goroutines antes de fechar o programa.
*/

func Exercicio4() {
	canal := make(chan string, 10)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			canal <- fmt.Sprintf("Mensagem %d", i)
		}
		close(canal)
		waitGroup.Done()
	}()

	go func() {
		for msg := range canal {
			fmt.Println(msg)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
