package controle_fluxo

import "fmt"

func ControleFluxo() {
	// Exemplode if-else
	idade := 20
	if idade >= 18 {
		fmt.Println("Você é maiorde idade")
	} else {
		fmt.Println("Você é menor de idade")
	}

	// Exemplo de for loop
	fmt.Println("Contagem de 1 a 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exemplo de switch
	diaDaSemana := 3
	fmt.Print("Hoje é ")
	switch diaDaSemana {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	}
}
