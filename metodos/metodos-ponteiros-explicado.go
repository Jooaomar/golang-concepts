package metodos

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func Abs(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MetodosPonteirosExplicado() {
	v := Vertex3{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
