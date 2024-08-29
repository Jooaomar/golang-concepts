package interfaces

import "fmt"

type I3 interface {
	M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func InterfacesValoresNil() {
	// Se o valor concreto no interior da própria interface é nulo, o método será chamado com um receptor nulo.

	// Em algumas linguagens isso iria desencadear uma exceção de ponteiro nulo, mas em Go é comum para escrever
	// métodos que normalmente lidam sendo chamado com um receptor nil (como acontece com o método M neste exemplo).

	// Nota-se que um valor de interface que contém um valor nulo concreto é por si não-nulo.

	var i I3

	var t *T3
	i = t // assume valor de ponteiro de T3
	describe2(i)
	i.M()

	i = &T3{"Hello"}
	describe2(i)
	i.M()
}

func describe2(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}
