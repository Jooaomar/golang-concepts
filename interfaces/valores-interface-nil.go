package interfaces

import "fmt"

type I4 interface {
	M()
}

func ValoresInterfaceNil() {
	// Um valor de interface nil detém nem o valor nem tipo concreto.

	// Chamar um método em uma interface nil resulta em um erro de tempo
	// de execução porque não há um tipo dentro da tupla de interface para
	// indicar qual o método concreto a chamar.

	var i I4
	describe3(i)
	i.M()
}

func describe3(i I4) {
	fmt.Printf("(%v, %T)\n", i, i)
}
