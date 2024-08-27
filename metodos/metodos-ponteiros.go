// Métodos com receptores ponteiro podem modificar o valor que o receptor aponta (como Scale faz aqui).
package metodos

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	// Mudança de valores definido na instáncia de Vertex2
	v.X = v.X * f
	v.Y = v.Y * f
}

func MetodosPonteiros() {
	v := Vertex2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
