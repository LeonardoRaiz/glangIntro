package main

import "fmt"

type user struct {
	nome     string
	idade    uint8
	endereco endereco
} // molde do struct

type endereco struct {
	logradouro string
	numero     int
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type estudante struct {
	pessoa    // Quase herança xD
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Lana",
		"sobrenome": "Raiz",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}

	fmt.Println(usuario2)

	fmt.Println("******************** Structs ******************")

	var u user
	u.nome = "Leonardo"
	u.idade = 34
	fmt.Println(u)
	fmt.Println(u.nome)
	fmt.Println(u.idade)

	usuario3 := user{"Carolina", 34, endereco{logradouro: "Cruz"}}
	fmt.Println(usuario3)

	usuario4 := user{idade: 10}
	fmt.Println(usuario4)

	fmt.Println("***************** Herança... Só que não hehe **********************")
	pessoa1 := pessoa{"José", "Carlos", 20}
	fmt.Println(pessoa1)
	estudandte1 := estudante{pessoa1, "ADS", "FATEC"}
	fmt.Println(estudandte1)
	fmt.Println(estudandte1.nome)

}
