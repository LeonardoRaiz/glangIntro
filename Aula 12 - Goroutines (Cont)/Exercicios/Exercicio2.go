package Exercicios

import (
	"fmt"
	"sync"
	"time"
)

/*
Exercício 2: WaitGroup para Múltiplas Goroutines
Expanda o uso de WaitGroup para sincronizar cinco goroutines realizando tarefas diferentes. Cada tarefa deve imprimir uma série de números de 1 a 5, com cada número sendo impresso após um segundo. A main goroutine deve apenas terminar após todas as outras terem completado.
*/

func Exercicio2() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	for i := 0; i <= 5; i++ {
		go func(id int) {
			for j := 0; j <= 5; j++ {
				fmt.Printf("Goroutine %d: %d\n", id, j)
				time.Sleep(time.Second)
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}
