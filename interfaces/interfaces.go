package interfaces

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Este método significa que o tipo T implementa a interface I,
// mas não precisamos declarar explicitamente que isso acontece.
func (t T) M() {
	fmt.Println(t.S)
}

func Interface() {
	var i I = T{"Hello"} // definimos a interface é como se dissesemos: "T precisa ter um método 'M()' "
	i.M()
}
