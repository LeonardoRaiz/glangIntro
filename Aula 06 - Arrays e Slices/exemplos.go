package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	for i := 0; i < len(array1); i++ {
		fmt.Println("Informe o valor do array na posição [", i, "]")
		fmt.Scanln(&array1[i])
	}
	fmt.Println(array1)

	// Não é tamanho dinâmico
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// Slice é dinâmico, ou seja, ele aponta para um array (fatia do array)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 11) // adicionar valores no slice
	fmt.Println(slice)

	slice2 := array2[1:3] //inclusivo exclusivo
	fmt.Println(slice2)
	array2[1] = 1000
	//Slice alterado pelo ponteiro
	fmt.Println(slice2)

	fmt.Println("**************** Array interno *****************")

	slice3 := make([]float32, 10, 11) // tipo, tamanho do slice, capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade
	//e se estourar a capacidade?

	slice3 = append(slice3, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	//GO dobra o tamanho do slice
	slice3 = append(slice3, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

}
