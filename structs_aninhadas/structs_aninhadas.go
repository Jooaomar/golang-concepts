package structs_aninhadas

import "fmt"

// Definindo uma struct Endereco
type Endereco struct {
	Rua    string
	Cidade string
	CEP    string
}

// Definindoums struct Pessoa com um campo Endereco
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func ExibirStructsAninhadas() {
	// Criando uma instáncia da struct Pessoa
	p := Pessoa{
		Nome:  "Alice",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua Exemplo, 123",
			Cidade: "Cidade Exemplo",
			CEP:    "12345-678",
		},
	}

	// Acessando campos aninhados
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)
	fmt.Println("Endereço:", p.Endereco.Rua, p.Endereco.Cidade, p.Endereco.CEP)
}
