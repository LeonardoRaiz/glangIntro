package main

import (
	"Introducao/enderecos"
	"fmt"
)

func main() {
	tipoEnderecos := enderecos.TipoDeEndereco("Rodovia Paulista")
	fmt.Println(tipoEnderecos)
}
