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

func (p Product) Describe() (string, decimal.Decimal) {
	total := p.CalculateTotal()
	desc := fmt.Sprintf("%v kg %v cost %v USD!", p.Quantity, p.Name, total)
	return desc, total
}

func main() {
	products := []Product{
		{
			Name:       "Watermelon",
			PricePerKg: decimal.NewFromFloat(1.00),
			Quantity:   2.5,
		},
		{
			Name:       "Banana",
			PricePerKg: decimal.NewFromFloat(2.00),
			Quantity:   1.5,
		},
	}
	//products := []Product{watermelon, banana}

	grandTotal := decimal.Zero
	for _, p := range products {
		desc, total := p.Describe()
		fmt.Println(desc)
		grandTotal = grandTotal.Add(total)
	}
	fmt.Printf("Grand Total: %v USD\n", grandTotal)
}
