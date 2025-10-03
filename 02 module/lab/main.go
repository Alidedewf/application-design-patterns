package main

import "fmt"

// DRY
type OrderService struct{}

func (o *OrderService) calculateTotal(quantity int, price float64) float64 {
	return float64(quantity) * price
}

func (o *OrderService) CreateOrder(productName string, quantity int, price float64) {
	total := o.calculateTotal(quantity, price)
	fmt.Printf("Заказ на %s создан. Итого: %.2f\n", productName, total)
}

func (o *OrderService) UpdateOrder(productName string, quantity int, price float64) {
	total := o.calculateTotal(quantity, price)
	fmt.Printf("Заказ для %s обновлён. Новая сумма: %.2f\n", productName, total)
}

type Vehicle struct {
	Type string
}

func (v *Vehicle) Start() {
	fmt.Printf("%s запускается\n", v.Type)
}

func (v *Vehicle) Stop() {
	fmt.Printf("%s останавливается\n", v.Type)
}

// KISS
func Add(a, b int) {
	fmt.Println("Sum = ", a+b)
}

type Singleton struct{}

func (s *Singleton) DoSomething() {
}

var instance = &Singleton{}

// YAGNI
func CircleArea(radius float64) float64 {
	return 3.14 * radius * radius
}

func SquareArea(side float64) float64 {
	return side * side
}

func SimpleAdd(a, b int) int {
	return a + b
}

func main() {
	// DRY
	order := OrderService{}
	order.CreateOrder("Book", 2, 500)
	order.UpdateOrder("Book", 3, 500)

	car := Vehicle{Type: "Car"}
	car.Start()
	car.Stop()

	// KISS
	Add(10, 20)
	instance.DoSomething()

	// YAGNI
	fmt.Println("Площадь круга:", CircleArea(5))
	fmt.Println("Площадь площади:", SquareArea(4))
	fmt.Println("Простое дополнение:", SimpleAdd(7, 8))
}
