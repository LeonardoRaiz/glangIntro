package enderecos

import "strings"

// TipoDeEndereco de endereço verifica se tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"rua",
		"avenida",
		"estrada",
		"rodovia",
	}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	// RUA DOS TESTES
	// ["RUA", "DOS", "TESTES"]

	enderecoTemTipoValido := false
	for _, tipos := range tiposValidos {
		if tipos == primeiraPalavraEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {

		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}
