package Exercicios

import (
	"fmt"
	"time"
)

/*
Exercício 3: Uso de Canais com Range
Modifique a função escreverChannel para enviar uma sequência de números (de 1 a 10) em vez de uma mensagem repetida. Na função main, utilize um loop range para ler e imprimir esses números do canal. Observe como o canal transmite os dados de uma goroutine para outra.
*/

func Exercicio3() {
	canal := make(chan int)
	go escreverNumeros(canal)
	for numero := range canal {
		fmt.Println(numero)
		time.Sleep(time.Second)
	}
}

func escreverNumeros(canal chan<- int) {
	for i := 1; i <= 10; i++ {
		canal <- i
	}
	close(canal)
}
