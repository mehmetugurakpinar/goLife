package main

import "fmt"

type fibonacci struct {
	a, b int
}

func (f *fibonacci) increment() int {
	r := f.a
	f.a, f.b = f.b, f.a+f.b
	return r
}

func main() {
	f := &fibonacci{b: 1}
	for range 10 {
		fmt.Println(f.increment())
	}
}
