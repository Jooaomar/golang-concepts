package sintaxe

import "fmt"

// A declaração return sem argumentos retorna os valores atuais dos resultados. Isto é conhecido como um retorno "limpo".
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func RetornoLimpo() {
	fmt.Println(split(17))
}
