package Exercicios

import "fmt"

/*
Exercício 5: Deadlock Prevention
Desenvolva um programa que ilustra um caso de deadlock potencial e então modifique-o para prevenir o deadlock. O programa deve usar pelo menos dois canais e duas goroutines que comunicam entre si, onde um mal-entendido na ordem de operações poderia causar o deadlock.
*/

func Exercicio5() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		canal1 <- "Enviado para canal 1"
		resp := <-canal2
		fmt.Println("Recebido no canal 1:", resp)
	}()

	go func() {
		canal2 <- "Enviado para canal 2"
		resp := <-canal1
		fmt.Println("Recebido no canal 2:", resp)
	}()

	fmt.Scanln() // Espera input para não terminar o programa imediatamente
}
