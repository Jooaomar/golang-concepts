package interfaces

import (
	"fmt"
	"math"
)

type I2 interface {
	M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func InterfacesValores() {
	var i I2
	i = &T2{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
