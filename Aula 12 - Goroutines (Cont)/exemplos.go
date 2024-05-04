package main

import (
	"fmt"
	"time"
)

func main() {
	//Canais com Buffer
	canal := make(chan string, 2)
	canal <- "Olá Mundo"
	canal <- "Programando em Go"

	msg := <-canal
	fmt.Println(msg)

	//SELECT para concorrência
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			canal2 <- "Canal 2"
		}
	}()

	for {

		select {
		case msgCanal1 := <-canal1:
			fmt.Println(msgCanal1)

		case msgCanal2 := <-canal2:
			fmt.Println(msgCanal2)
		}
		//msgCanal1 := <-canal1
		//fmt.Println(msgCanal1)
		//
		//msgCanal2 := <-canal2
		//fmt.Println(msgCanal2)
	}
}
