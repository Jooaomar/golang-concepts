package estruturas_dados

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

func MapsLiterais() {
	fmt.Println(m2)
}
