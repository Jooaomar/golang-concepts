// Entendendo
package metodos

import "fmt"

type Vertex4 struct {
	X, Y float64
}

// métodos com receptores ponteiro assumem um valor ou um ponteiro
// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// funções com um argumento de ponteiro deve receber somente ponteiro
// var v Vertex
// ScaleFunc(v, 5)  // Erro de compilação!
// ScaleFunc(&v, 5) // OK
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func IndirecaoPonteiros() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
