package goroutines

import (
	"fmt"
	"time"
)

func Exibir() {
	// Funçãoque será executadacomo goroutine
	go sayHello()

	// Aguarda um pouco para garantir que a goroutine seja executada
	time.Sleep(1 * time.Second)
}

func sayHello() {
	fmt.Println("Hello, John Marcel")
}
