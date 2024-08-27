package metodos

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	// Isso -> '(v Vertex)',  se chama RECEPTOR que qualifica Abs() como um método pertencente a Vertex
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Metodos() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
