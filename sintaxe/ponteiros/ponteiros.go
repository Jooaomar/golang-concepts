package ponteiros

import "fmt"

func Ponteiros() {
	i, j := 42, 2701

	p := &i         // ponit to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i throught the pointer
	fmt.Println(i)  // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide  j through
	fmt.Println(j)
}
