// Em Go, nil é um valor especial que representa a ausência de valor para tipos de dados
// que são referências, como ponteiros, slices, mapas, funções, interfaces e canais.
// Quando uma variável desses tipos é declarada e não inicializada, seu valor padrão é nil,
// o que significa que ela não aponta para nenhum valor ou objeto.
package tiposbasicos

import "fmt"

func NilTipo() {
	// exemplo com ponteiro
	var p *int
	fmt.Println("Ponteiro p:", p)

	if p == nil {
		fmt.Println("O ponteiro p é nil")
	}

	// Exemplo com slice
	var s []int
	fmt.Println("Slice s:", s)              // Saída: Slice s: []
	fmt.Println("Slice s é nil:", s == nil) // Saída: Slice s é nil: true

	if s == nil {
		fmt.Println("O slice s é nil")
	}

	// Exemplo com mapa
	var m map[string]int
	fmt.Println("Mapa m:", m)              // Saída: Mapa m: map[]
	fmt.Println("Mapa m é nil:", m == nil) // Saída: Mapa m é nil: true

	if m == nil {
		fmt.Println("O mapa m é nil")
		m = make(map[string]int) // Inicializando o mapa
		m["chave"] = 42
		fmt.Println("Mapa m após inicialização:", m)
	}

	// Exemplo com canal
	var c chan int
	fmt.Println("Canal c:", c)              // Saída: Canal c: <nil>
	fmt.Println("Canal c é nil:", c == nil) // Saída: Canal c é nil: true

	if c == nil {
		fmt.Println("O canal c é nil")
		c = make(chan int) // Inicializando o canal
		go func() {
			c <- 42
		}()
		fmt.Println("Valor recebido do canal c:", <-c)
	}

	// Exemplo com interface
	var i interface{}
	fmt.Println("Interface i:", i)              // Saída: Interface i: <nil>
	fmt.Println("Interface i é nil:", i == nil) // Saída: Interface i é nil: true

	if i == nil {
		fmt.Println("A interface i é nil")
		i = "Agora i não é nil"
		fmt.Println("Interface i após atribuição:", i)
	}
}
