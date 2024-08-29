package interfaces

import "fmt"

func TypeAssertions() {
	// A type assertion fornece acesso ao valor concreto subjacente de um valor de interface.
	// Esta declaração afirma que o valor de interface i detém o tipo concreto T e atribui o valor
	// de T subjacente à variável t.
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// Para testar se um valor de interface possui um tipo específico, uma type assertion_ pode retornar
	// dois valores: o valor subjacente e um valor booleano que informa se a afirmação é correta.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Se i não detém uma T, a declaração irá desencadear um erro pânico.
	f = i.(float64) // panic: interface conversion: interface {} is string, not float64
	fmt.Println(f)
}
