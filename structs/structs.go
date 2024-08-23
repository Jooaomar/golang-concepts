package structs

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

func ExibirStructs() {
	// Criando uma instáncia da struct Pessoa
	p := Pessoa{
		Nome:   "Alice",
		Idade:  30,
		Altura: 1.75,
	}

	// Acessandoe exibindo
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)
	fmt.Println("Altura:", p.Altura)
}
