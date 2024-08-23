// Channels podem ser "buffered" (com buffer) ou "unbuffered" (sem buffer). Um channel sem buffer bloqueia
// tanto o envio quanto o recebimento até que ambos estejam prontos. Um channel com buffer permite enviar
// vários valores antes que algum deles seja recebido.

package goroutines

import "fmt"

func ChamaBuffered() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// Recebe e imprmie os valores. A notação "<-" retira os valores armazenados no channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
