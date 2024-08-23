package structs_metodos

import "fmt"

type Retangulo struct {
	Largura float64
	Altura  float64
}

// Método para calcular a área
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Mátodo para calcular o perímetro
func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

func ExibirStructsMetodos() {
	// Criando uma instáncia da struct Retangulo
	r := Retangulo{
		Largura: 4.0,
		Altura:  5.0,
	}

	// Chamando métodos da struct
	fmt.Println("Áreado retângulo:", r.Area())
	fmt.Println("Perímetro do retângulo:", r.Perimetro())
}
