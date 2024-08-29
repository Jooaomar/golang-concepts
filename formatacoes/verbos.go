package formatacoes

import "fmt"

func Verbos() {
	// Exemplo de %s - String
	nome := "Alice"
	fmt.Printf("Nome: %s\n", nome) // Saída: Nome: Alice

	// Exemplo de %T - Tipo
	idade := 30
	fmt.Printf("Tipo de 'idade': %T\n", idade) // Saída: Tipo de 'idade': int

	// Exemplo de %v - Valor padrão
	empregado := true
	fmt.Printf("Empregado: %v\n", empregado) // Saída: Empregado: true

	// Exemplo de %d - Número inteiro decimal
	fmt.Printf("Idade: %d\n", idade) // Saída: Idade: 30

	// Exemplo de %f - Número de ponto flutuante
	salario := 45000.50
	fmt.Printf("Salário: %f\n", salario) // Saída: Salário: 45000.500000

	// Exemplo de %g - Número de ponto flutuante (notação científica ou decimal)
	pi := 3.14159265359
	fmt.Printf("Pi: %g\n", pi) // Saída: Pi: 3.14159

	// Exemplo de %b - Número binário
	numero := 7
	fmt.Printf("Binário de 7: %b\n", numero) // Saída: Binário de 7: 111

	// Exemplo de %x - Número hexadecimal (letras minúsculas)
	fmt.Printf("Hexadecimal de 255: %x\n", 255) // Saída: Hexadecimal de 255: ff

	// Exemplo de %X - Número hexadecimal (letras maiúsculas)
	fmt.Printf("Hexadecimal de 255: %X\n", 255) // Saída: Hexadecimal de 255: FF

	// Exemplo de %p - Endereço de memória
	p := &idade
	fmt.Printf("Endereço de memória de 'idade': %p\n", p) // Exemplo de saída: Endereço de memória de 'idade': 0xc0000160d8

	// Exemplo de %q - Exibição de Valores Literais
	cores := []string{"Vermelho", "Verde", "Azul"}
	fmt.Printf("As cores são: %q\n", cores) // As cores são: ["Vermelho" "Verde" "Azul"]

}
