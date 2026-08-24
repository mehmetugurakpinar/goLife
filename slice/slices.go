package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11}

	var s []int = primes[1:4]
	addItem(s, 13)
	s = addItemWithReturn(s, 15)
	fmt.Println(s)
}

func addItem(s []int, item int) {
	s = append(s, item)
	fmt.Println(s)
}

func addItemWithReturn(s []int, item int) []int {
	s = append(s, item)
	fmt.Println(s)
	return s
}
