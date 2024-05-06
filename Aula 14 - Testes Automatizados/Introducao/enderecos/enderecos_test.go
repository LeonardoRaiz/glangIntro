package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//TESTE DE UNIDADE
//Testa uma pequena parte do código

func TestTipoDeEndereco(t *testing.T) {
	// Test com T

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS TESTES", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{" ", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}

	//enderecoParaTeste := "Rua ABC"
	//tipoDeEnderecoEsperado := "Avenida"
	//tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

}
