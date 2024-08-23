package constantes

import "fmt"

const (
	StatusOK       = 200
	StatusNotFound = 404
)

func EvitaMagicaDeCodigo() {
	const statusCode = 200
	// Em vez de:
	if statusCode == 200 {
		// faz algo
		fmt.Println("Ok 1")
	}
	// Use:
	if statusCode == StatusOK {
		// faz algo
		fmt.Println("Ok 2")
	}
}
