package switchfluxo

import (
	"fmt"
	"time"
)

func SwitchSemCondicao() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("Good evening.")
	}
}
