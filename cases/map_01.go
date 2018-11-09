package cases

import (
	"fmt"
)

type Vertex struct {
	Lat  float64
	Long float64
}

type MyMap map[interface{}]interface{}

func Map01() {
	var m map[string]Vertex
	m = map[string]Vertex{}
	m["abc"] = Vertex{123.12, 432.12}
	m["bcd"] = Vertex{45.1, -123.343}

	fmt.Printf("%T %v %v\n", m, m, len(m))

	n := MyMap{
		"Tom":   Vertex{1, 2},
		"Jack":  Vertex{3, 4},
		"Amose": "345234",
		123:     456,
	}
	fmt.Printf("%T %v %v\n", n, n, len(n))

	elem, ok := n["tom"]
	if ok {
		fmt.Println(elem)
	}
}
