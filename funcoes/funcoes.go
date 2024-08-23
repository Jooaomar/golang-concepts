package funcoes

import "fmt"

func saudacao(nome string) {
	fmt.Println("Olá,", nome)
}

func somar(a int, b int) int {
	return a + b
}

func dividir(a float64, b float64) (float64, float64) {
	quociente := a / b
	resto := a - (b * float64(int(a/b)))
	return quociente, resto
}

func imprimirNumeros(numeros ...int) {
	fmt.Println("Números fornecidos")
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}

func ExibirFuncoes() {
	saudacao("Alice")

	resultado := somar(10, 20)
	fmt.Println("Resultado da soma: ", resultado)

	quocient, resto := dividir(10, 3)
	fmt.Printf("Quociente: %.2f, Resto: %.2f\n", quocient, resto)

	imprimirNumeros(1, 2, 3, 4, 5)
}
