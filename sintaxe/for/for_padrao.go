package forfluxo

import "fmt"

func ForPadrao() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
