package variaveis

import "fmt"

func ExibirVariaveis() {
	// Declaração explícita com 'var' você pode especificar o tipo
	var nome string = "Alice"
	var idade int = 30
	var salario float64 = 45000.50
	var empregado bool = true
	// Poderia também:
	// var idade int // idade será inicializada com 0
	// var nome string // nome será inicializado com ""

	// Declaração implícita com ':=' sem a necessidade de explicitar o tipo
	// OBS: Não é possível usar := sem fornecer um valor inicial.
	sobrenome := "Silva"
	anosExperiencia := 5
	altura := 1.75
	freelancer := false

	// Exibindo valores e tipos
	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)
	fmt.Println("Idade:", idade)
	fmt.Println("Anos de Experiência:", anosExperiencia)
	fmt.Println("Salário:", salario)
	fmt.Println("Altura:", altura)
	fmt.Println("Empregado:", empregado)
	fmt.Println("Freelancer:", freelancer)

	// Exibindo tipos de dados
	fmt.Printf("Tipo de 'nome': %T\n", nome)
	fmt.Printf("Tipo de 'idade': %T\n", idade)
	fmt.Printf("Tipo de 'salario': %T\n", salario)
	fmt.Printf("Tipo de 'empregado': %T\n", empregado)
}
