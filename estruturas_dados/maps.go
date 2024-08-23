package estruturas_dados

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func MapBasic() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
