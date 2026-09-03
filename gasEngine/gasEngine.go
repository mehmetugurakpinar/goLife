package main

import "fmt"

type gasEngine struct {
	mpg     int
	gallons int
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

type hydrogenEngine struct {
	mpkg int
	kg   int
}

func (e gasEngine) milesLeft() int {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() int {
	return e.kwh * e.mpkwh
}

func (e hydrogenEngine) milesLeft() int {
	return e.kg * e.mpkg
}

type engine interface {
	milesLeft() int
}

func canMakeIt(e engine, miles int) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	myEngine := gasEngine{25, 15}
	myEngineElectric := electricEngine{25, 150}
	myHydrogenEngine := hydrogenEngine{25, 150}
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())
	fmt.Printf("Total miles left in battery: %v\n", myEngineElectric.milesLeft())
	canMakeIt(myEngine, 100)
	canMakeIt(myEngineElectric, 3300)
	canMakeIt(myHydrogenEngine, 3300)
}
