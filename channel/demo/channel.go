package main

import (
	"fmt"
)

func main() {
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
		//time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c)
	for i := range 5 {
		//if i == 4 {
		//	time.Sleep(2000 * time.Millisecond)
		//}
		c <- i
	}
	fmt.Println("Existing process...")
}
