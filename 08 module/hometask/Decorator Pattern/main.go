package main

import (
	"fmt"
)

type Beverage interface {
	Description() string
	Cost() float64
}

type Americano struct{}

func (e Americano) Description() string { return "Американо" }
func (e Americano) Cost() float64       { return 3.00 }

type Tea struct{}

func (t Tea) Description() string { return "Чай" }
func (t Tea) Cost() float64       { return 2.50 }

type Latte struct{}

func (l Latte) Description() string { return "Латте" }
func (l Latte) Cost() float64       { return 4.00 }

type Capuccino struct{}

func (m Capuccino) Description() string { return "Капучино" }
func (m Capuccino) Cost() float64       { return 4.50 }

// ====== Decorator Base ======
type BeverageDecorator struct {
	beverage Beverage
}

func (d BeverageDecorator) Description() string {
	return d.beverage.Description()
}

func (d BeverageDecorator) Cost() float64 {
	return d.beverage.Cost()
}

type Milk struct {
	BeverageDecorator
}

func NewMilk(b Beverage) Beverage {
	return &Milk{BeverageDecorator{b}}
}

func (m *Milk) Description() string { return m.beverage.Description() + ", Молоко" }
func (m *Milk) Cost() float64       { return m.beverage.Cost() + 0.50 }

type Sugar struct {
	BeverageDecorator
}

func NewSugar(b Beverage) Beverage {
	return &Sugar{BeverageDecorator{b}}
}

func (s *Sugar) Description() string { return s.beverage.Description() + ", Сахар" }
func (s *Sugar) Cost() float64       { return s.beverage.Cost() + 0.25 }

type WhippedCream struct {
	BeverageDecorator
}

func NewWhippedCream(b Beverage) Beverage {
	return &WhippedCream{BeverageDecorator{b}}
}

func (w *WhippedCream) Description() string { return w.beverage.Description() + ", Сливки" }
func (w *WhippedCream) Cost() float64       { return w.beverage.Cost() + 0.75 }

type Syrup struct {
	BeverageDecorator
}

func NewSyrup(b Beverage) Beverage {
	return &Syrup{BeverageDecorator{b}}
}

func (s *Syrup) Description() string { return s.beverage.Description() + ", Сироп" }
func (s *Syrup) Cost() float64       { return s.beverage.Cost() + 0.60 }

func testDecorator() {
	var drink Beverage = Americano{}
	drink = NewMilk(drink)
	drink = NewSugar(drink)
	drink = NewWhippedCream(drink)

	fmt.Println("Order:", drink.Description())
	fmt.Printf("Total: $%.2f\n", drink.Cost())

	var drink2 Beverage = Latte{}
	drink2 = NewSyrup(drink2)
	drink2 = NewSugar(drink2)

	fmt.Println("Order:", drink2.Description())
	fmt.Printf("Total: $%.2f\n", drink2.Cost())
}

func main() {
	fmt.Println("=== Decorator Pattern Demo ===")
	testDecorator()
}
