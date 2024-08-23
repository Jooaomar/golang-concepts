// Em Go, as constantes são valores fixos que são conhecidos em tempo de compilação e não
// podem ser alterados durante a execução do programa.
package sintaxe

import "fmt"

const Pi = 3.14

// Constantes são declaradas como variáveis, mas com a palavra-chave const.
// Constantes podem ser seqüências de caracteres, booleanos, ou valores numéricos.
// Constantes não podem ser declaradas usando a sintaxe :=.
func Constates() {
	const World = "João Marcelo"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
