package Exercicios

import (
	"fmt"
	"math/rand"
	"time"
)

func Exercicio3() {
	rand.Seed(time.Now().UnixNano())
	numeroAleatorio := rand.Intn(10) + 1
	tentativasMaximas := 3

	for tentativa := 1; tentativa <= tentativasMaximas; tentativa++ {
		var palpite int
		fmt.Print("Adivinhe o número (entre 1 e 10): ")
		fmt.Scanln(&palpite)

		if palpite == numeroAleatorio {
			fmt.Println("Parabéns! Você acertou.")
			break
		} else {
			fmt.Println("Tente novamente.")
		}

		if tentativa == tentativasMaximas {
			fmt.Printf("O número correto era %d.\n", numeroAleatorio)
		}
	}
}
