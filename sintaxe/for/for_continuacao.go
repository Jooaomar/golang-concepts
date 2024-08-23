package forfluxo

import "fmt"

func ForContinuacao() {
	sum := 1
	for sum < 1000 {
		sum += 1
	}
	fmt.Println(sum)
}
