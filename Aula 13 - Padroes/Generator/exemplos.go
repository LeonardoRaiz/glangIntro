package main

import (
	"fmt"
	"time"
)

func main() {
	// Generator para esconder a complexidade
	canal := escrever("Olá mundo")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
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
