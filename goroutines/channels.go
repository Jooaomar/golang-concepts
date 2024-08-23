// Channels são usados para comunicação segura entre goroutines. Eles permitem que uma goroutine envie dados para outra goroutine.
package goroutines

import (
	"fmt"
	"time"
)

func ChamaChannels() {
	// Criação de um channel
	messages := make(chan string)

	// Executa um goroutine para enviar uma mensagem
	go func() {
		messages <- "Hello, João" // Envia uma mensagem para o channel
	}()

	// Recebe a mensagem do Channel
	msg := <-messages
	fmt.Println(msg)

	time.Sleep(1 * time.Second) // Para evitar a saída antes da execução da goroutine
}
