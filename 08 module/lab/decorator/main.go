package main

import (
	"fmt"
)

type IBeverage interface {
	GetCost() float64
	GetDescription() string
}

// Coffee
type Coffee struct{}

func (c *Coffee) GetCost() float64 {
	return 750.0 
}

func (c *Coffee) GetDescription() string {
	return "Кофе"
}

// MilkDecorator
type MilkDecorator struct {
	beverage IBeverage 
}
func NewMilkDecorator(b IBeverage) *MilkDecorator {
	return &MilkDecorator{beverage: b}
}

func (m *MilkDecorator) GetCost() float64 {
	return m.beverage.GetCost() + 50.0 
}

func (m *MilkDecorator) GetDescription() string {
	return m.beverage.GetDescription() + ", Молоко"
}

// SugarDecorator
type SugarDecorator struct {
	beverage IBeverage
}


func NewSugarDecorator(b IBeverage) *SugarDecorator {
	return &SugarDecorator{beverage: b}
}

func (s *SugarDecorator) GetCost() float64 {
	return s.beverage.GetCost() + 5.0 
}

func (s *SugarDecorator) GetDescription() string {
	return s.beverage.GetDescription() + ", Сахар"
}

// ChocolateDecorator
type ChocolateDecorator struct {
	beverage IBeverage
}

func NewChocolateDecorator(b IBeverage) *ChocolateDecorator {
	return &ChocolateDecorator{beverage: b}
}

func (c *ChocolateDecorator) GetCost() float64 {
	return c.beverage.GetCost() + 50.0 
}

func (c *ChocolateDecorator) GetDescription() string {
	return c.beverage.GetDescription() + ", Шоколад"
}

// 3. Доп.декоратор

// VanillaDecorator
type VanillaDecorator struct {
	beverage IBeverage
}

func NewVanillaDecorator(b IBeverage) *VanillaDecorator {
	return &VanillaDecorator{beverage: b}
}

func (v *VanillaDecorator) GetCost() float64 {
	return v.beverage.GetCost() + 70.0
}

func (v *VanillaDecorator) GetDescription() string {
	return v.beverage.GetDescription() + ", Ваниль"
}

// CinnamonDecorator
type CinnamonDecorator struct {
	beverage IBeverage
}

func NewCinnamonDecorator(b IBeverage) *CinnamonDecorator {
	return &CinnamonDecorator{beverage: b}
}

func (c *CinnamonDecorator) GetCost() float64 {
	return c.beverage.GetCost() + 5.0 
}

func (c *CinnamonDecorator) GetDescription() string {
	return c.beverage.GetDescription() + ", Корица"
}

func main() {
	fmt.Println("Паттерн Декоратор")
	fmt.Println()

	var beverage1 IBeverage = &Coffee{}
	fmt.Printf("Заказ 1: %s \nИтого: %.2f тенге.\n", beverage1.GetDescription(), beverage1.GetCost())
	fmt.Println("--------------------")

	// Сцен.2
	var beverage2 IBeverage = &Coffee{}
	beverage2 = NewMilkDecorator(beverage2)   
	beverage2 = NewSugarDecorator(beverage2) 
	fmt.Printf("Заказ 2: %s \nИтого: %.2f тенге.\n", beverage2.GetDescription(), beverage2.GetCost())
	fmt.Println("--------------------")

	// Сцен.3
	var beverage3 IBeverage = &Coffee{}
	beverage3 = NewChocolateDecorator(beverage3) 
	beverage3 = NewVanillaDecorator(beverage3)   
	beverage3 = NewCinnamonDecorator(beverage3)  
	beverage3 = NewMilkDecorator(beverage3)  
	fmt.Printf("Заказ 3: %s \nИтого: %.2f тенге.\n", beverage3.GetDescription(), beverage3.GetCost())
}