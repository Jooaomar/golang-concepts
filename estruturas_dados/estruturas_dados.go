package estruturas_dados

import (
	"fmt"
)

// Struct para representar um Pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

// Função para mostrar informações de um slice de Pessoas
func mostrarPessoas(pessoas []Pessoa) {
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)
	}
}

func ExibirEstruturaDeDados() {
	// arrays
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println("Array de números; ", numeros)

	// Slices
	cores := []string{"Vermelho", "Verde", "Azul"}
	cores = append(cores, "Amarelo")
	fmt.Println("Slice de cores: ", cores)

	// Criando e inicializando um slice de inteiros
	sliceNumero := []int{1, 2, 3, 4, 5}

	// Exibindo o slice
	fmt.Println("Slice de números inteiros:", sliceNumero)

	// Maps
	capitais := make(map[string]string)
	capitais["Brasil"] = "Brasília"
	capitais["França"] = "Paris"
	capitais["Japão"] = "Tóquio"

	fmt.Println("Mapa de capitais: ")
	for pais, capital := range capitais {
		fmt.Printf("Pais: %s, Capital %s\n", pais, capital)
	}

	// Structs
	pessoas := []Pessoa{
		{Nome: "Alice"},
		{Nome: "Bob", Idade: 25},
		{Nome: "Carol", Idade: 27},
	}

	fmt.Println("Lista de pessoas: ")
	mostrarPessoas(pessoas)
}
