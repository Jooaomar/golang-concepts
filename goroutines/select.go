// A palavra-chave select é usada para trabalhar com múltiplos channels ao mesmo tempo.
// Ela permite esperar em múltiplas operações de channels, e processa a primeira que estiver pronta.

package goroutines

import (
	"fmt"
	"time"
)

func ChamarSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second) // simulando processamento que leva tempo de resposta
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
		// As expressões <-c1 e <-c2 são operações de recepção de canal em Go.
		// <-c1: Recebe um valor do canal c1. Quando o código chega a esse ponto,
		// ele espera até que um valor esteja disponível no canal c1. Quando um valor é
		// enviado para c1, a expressão <-c1 captura esse valor, que é então atribuído à variável msg1.

		// Como funciona o select no seu código
		// O select bloqueia e espera até que um valor seja recebido de c1 ou c2.
		// Se c1 estiver pronto para enviar um valor, o valor será recebido, atribuído a msg1, e "Received", msg1 será impresso.
	}
}
