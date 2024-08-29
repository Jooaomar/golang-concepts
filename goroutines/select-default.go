package goroutines

import (
	"fmt"
	"time"
)

func SelectDefault() {
	tick := time.Tick(100 * time.Millisecond)
	bom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-bom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
