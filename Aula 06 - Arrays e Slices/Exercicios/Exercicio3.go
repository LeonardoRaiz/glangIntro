package Exercicios

import "fmt"

/*
Exercício 3: Fatias de Arrays
Utilize o array [...]int{5, 10, 15, 20, 25, 30} e crie três slices diferentes a partir dele:

Um que inclua os três primeiros elementos.
Um que inclua os últimos três elementos.
Um que inclua os elementos do meio.
Após criar cada slice, modifique o valor original do array que afeta o segundo elemento de cada slice criado. Imprima o array e os slices após as modificações para observar o impacto.
*/

func Exercicio3() {
	array := [...]int{5, 10, 15, 20, 25, 30}

	slice := array[0:3]
	slice2 := array[3:6]
	slice3 := array[1:5]

	fmt.Println("Antes")
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)

	array[1] = 100

	fmt.Println("Depois")
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
