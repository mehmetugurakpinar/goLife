package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Product struct {
	Name       string
	Quantity   float64
	PricePerKg decimal.Decimal
}

func (p Product) CalculateTotal() decimal.Decimal {
	return p.PricePerKg.Mul(decimal.NewFromFloat(p.Quantity))
}

func main() {
	watermelon := Product{
		Name:       "Watermelon",
		PricePerKg: decimal.NewFromFloat(1.00),
		Quantity:   2.5,
	}

	banana := Product{
		Name:       "Banana",
		PricePerKg: decimal.NewFromFloat(2.00),
		Quantity:   1.5,
	}

	watermelonTotal := watermelon.CalculateTotal()
	bananaTotal := banana.CalculateTotal()
	fmt.Printf("%v kg %v cost %v USD!\n", watermelon.Quantity, watermelon.Name, watermelonTotal)
	fmt.Printf("%v kg %v cost %v USD!\n", banana.Quantity, banana.Name, bananaTotal)
	fmt.Printf("Total price: %v USD\n", watermelonTotal.Add(bananaTotal))
}
