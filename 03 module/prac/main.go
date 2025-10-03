package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type IPayment interface {
	ProcessPayment(amount float64) error
}

type CreditCardPayment struct{}

func (c CreditCardPayment) ProcessPayment(amount float64) error {
	fmt.Println("Оплата кредитной картой на сумму:", amount)
	return nil
}

type PayPalPayment struct{}

func (p PayPalPayment) ProcessPayment(amount float64) error {
	fmt.Println("Оплата через PayPal на сумму:", amount)
	return nil
}

type IDelivery interface {
	DeliverOrder(order *Order) error
}

type CourierDelivery struct{}

func (c CourierDelivery) DeliverOrder(order *Order) error {
	fmt.Println("Доставка курьером для заказа:", order.ID)
	return nil
}

type PostDelivery struct{}

func (p PostDelivery) DeliverOrder(order *Order) error {
	fmt.Println("Доставка по почте для заказа:", order.ID)
	return nil
}

type INotification interface {
	SendNotification(message string)
}

type EmailNotification struct{}

func (e EmailNotification) SendNotification(message string) {
	fmt.Println("Email:", message)
}

type SmsNotification struct{}

func (s SmsNotification) SendNotification(message string) {
	fmt.Println("SMS:", message)
}

type IDiscount interface {
	CalculateDiscount(order *Order) float64
}

type FixedDiscount struct {
	Amount float64
}

func (f FixedDiscount) CalculateDiscount(order *Order) float64 {
	return f.Amount
}

type PercentDiscount struct {
	Percent float64
}

func (p PercentDiscount) CalculateDiscount(order *Order) float64 {
	return order.TotalWithoutDiscount() * p.Percent / 100
}

type Order struct {
	ID            int
	Products      []Product
	Payment       IPayment
	Delivery      IDelivery
	Notifications []INotification
	Discounts     []IDiscount
}

func (o *Order) AddProduct(p Product) {
	o.Products = append(o.Products, p)
}

func (o *Order) TotalWithoutDiscount() float64 {
	var total float64
	for _, p := range o.Products {
		total += p.Price * float64(p.Quantity)
	}
	return total
}

func (o *Order) TotalWithDiscount() float64 {
	total := o.TotalWithoutDiscount()
	for _, d := range o.Discounts {
		total -= d.CalculateDiscount(o)
	}
	return total
}

func (o *Order) ProcessOrder() {
	total := o.TotalWithDiscount()
	fmt.Println("Итоговая сумма заказа:", total)

	_ = o.Payment.ProcessPayment(total)
	_ = o.Delivery.DeliverOrder(o)

	for _, notifier := range o.Notifications {
		notifier.SendNotification(fmt.Sprintf("Ваш заказ #%d обработан. Сумма: %.2f", o.ID, total))
	}
}

func main() {
    order := &Order{
        ID: 1,
        Payment: CreditCardPayment{},
        Delivery: CourierDelivery{},
        Notifications: []INotification{EmailNotification{}, SmsNotification{}},
        Discounts: []IDiscount{
            FixedDiscount{Amount: 500},
            PercentDiscount{Percent: 10},
        },
    }

    order.AddProduct(Product{Name: "Ноутбук", Price: 250000, Quantity: 1})
    order.AddProduct(Product{Name: "Мышь", Price: 10000, Quantity: 2})

    order.ProcessOrder()
}