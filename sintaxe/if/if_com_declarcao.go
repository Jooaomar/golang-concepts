package iffluxo

import (
	"fmt"
	"math"
)

// Como o for, a instrução if pode começar com uma breve declaração antes de executar a condição.

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v // Variáveis ​​declaradas pela instrução são válidas somente no escopo até o final do if.
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // Variáveis ​​declaradas dentro de uma instrução if curta também estão disponíveis dentro dos blocos else.
	}
	// can't use v here, though
	return lim
}

func IfComDeclaracao() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
