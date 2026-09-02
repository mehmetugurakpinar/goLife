package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	createMap()

	m := map[string]int{"a": 1}
	fmt.Println(m)
	mutate(m)
	fmt.Println(m)
}

func createMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Kovanlik"] = Vertex{
		20.68433, -14.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Kovanlik"])
	clear(m)
	fmt.Println(m["Kovanlik"])
	counters := make(map[string]Vertex)
	counters["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	counters["Kovanlik"] = Vertex{
		20.68433, -14.39967,
	}
	fmt.Println(len(counters))
	fmt.Println(counters)
}

func mutate(m map[string]int) {
	m["b"] = 2
}
