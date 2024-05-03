package Exercicios

import "fmt"

/*
Exercício 4: Uso de Mapas Genéricos
Crie uma função que adiciona um conjunto de pares chave-valor a um mapa do tipo map[interface{}]interface{}. A função deve aceitar um mapa e um slice de chaves e um slice de valores (ambos de tipo []interface{}). Adicione valores ao mapa dentro da função e imprima o mapa após adicionar os itens.
*/

func addItems(m map[interface{}]interface{}, keys []interface{}, values []interface{}) {
	for i, key := range keys {
		m[key] = values[i]
	}
}

func Exercicio4() {
	mapa := make(map[interface{}]interface{})
	keys := []interface{}{"nome", "idade", 3}
	values := []interface{}{"John", 30, true}

	addItems(mapa, keys, values)

	for key, value := range mapa {
		fmt.Println("Chave:", key, "Valor:", value)
	}
}
