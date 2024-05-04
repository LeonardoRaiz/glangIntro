package main

import (
	"fmt"
	"time"
)

func main() {
	// juntar dois canais em um único canal
	canal := multiplaxar(escrever("Olá Mundo"), escrever("Programando em Go!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplaxar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case msg := <-canalEntrada1:
				canalSaida <- msg
			case msg := <-canalEntrada2:
				canalSaida <- msg

			}
		}
	}()
	return canalSaida
}

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %v", text)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return canal
}
