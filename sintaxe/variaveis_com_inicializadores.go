package sintaxe

import "fmt"

var i, j int = 1, 2

// A declaração var pode incluir inicializadores, uma por variável.
// Se um inicializador está presente, o tipo pode ser omitido; a variável terá o tipo do inicializador.
func VariaveisComInicializadores() {
	var c, python, php = true, false, "no!"
	fmt.Println(i, j, c, python, php)
}
