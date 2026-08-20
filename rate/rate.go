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

func newProduct(name string, quantity float64, PricePerKg decimal.Decimal) (Product, error) {
	if quantity <= 0 {
		return Product{}, fmt.Errorf("quantity must be positive")
	}
	if PricePerKg.IsZero() {
		return Product{}, fmt.Errorf("pricePerKg must be greater than zero")
	}

	return Product{name, quantity, PricePerKg}, nil
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
	apple, err := newProduct("Apple", 1.0, decimal.NewFromFloat(1.50))
	if err != nil {
		fmt.Println("Error creating apple:", err)
		return
	}
	orange, err := newProduct("Orange", 2.0, decimal.NewFromFloat(2.50))
	if err != nil {
		fmt.Println("Error creating orange:", err)
		return
	}
	products := []Product{apple, orange}

	grandTotal := decimal.Zero
	for _, p := range products {
		desc, total := p.Describe()
		fmt.Println(desc)
		grandTotal = grandTotal.Add(total)
	}
	fmt.Printf("Grand Total: %v USD\n", grandTotal)
}
