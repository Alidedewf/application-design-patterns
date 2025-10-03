package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

type Invoice struct {
	Id      int
	Items   []Item
	TaxRate float64
}

type InvoiceCalculator struct{}

func (c InvoiceCalculator) CalculateTotal(invoice Invoice) float64 {
	var subTotal float64
	for _, item := range invoice.Items {
		subTotal += item.Price
	}
	return subTotal + (subTotal * invoice.TaxRate)
}

type InvoiceRepository struct{}

func (r InvoiceRepository) Save(invoice Invoice) {
	fmt.Printf("Счет №%d сохранен в базу данных.\n", invoice.Id)
}

type DiscountStrategy interface {
	ApplyDiscount(amount float64) float64
}

type RegularCustomer struct{}

func (r RegularCustomer) ApplyDiscount(amount float64) float64 {
	return amount
}

type SilverCustomer struct{}

func (s SilverCustomer) ApplyDiscount(amount float64) float64 {
	return amount * 0.9
}

type GoldCustomer struct{}

func (g GoldCustomer) ApplyDiscount(amount float64) float64 {
	return amount * 0.8
}

type DiscountCalculator struct {
	strategy DiscountStrategy
}

func (dc DiscountCalculator) Calculate(amount float64) float64 {
	return dc.strategy.ApplyDiscount(amount)
}

type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

type Sleepable interface {
	Sleep()
}

type HumanWorker struct{}

func (h HumanWorker) Work() {
	fmt.Println("Человек работает")
}
func (h HumanWorker) Eat() {
	fmt.Println("Человек ест")
}
func (h HumanWorker) Sleep() {
	fmt.Println("Человек спит")
}

type RobotWorker struct{}

func (r RobotWorker) Work() {
	fmt.Println("Робот работает")
}

type Sender interface {
	Send(message string)
}

type EmailService struct{}

func (e EmailService) Send(message string) {
	fmt.Println("Email:", message)
}

type SmsService struct{}

func (s SmsService) Send(message string) {
	fmt.Println("SMS:", message)
}

type Notification struct {
	sender Sender
}

func (n Notification) SendNotification(message string) {
	n.sender.Send(message)
}

func main() {
	// === SRP ===
	invoice := Invoice{
		Id: 1,
		Items: []Item{
			{"Товар A", 100},
			{"Товар B", 200},
		},
		TaxRate: 0.2,
	}
	calculator := InvoiceCalculator{}
	total := calculator.CalculateTotal(invoice)
	fmt.Println("Итоговая сумма счета:", total)
	repo := InvoiceRepository{}
	repo.Save(invoice)

	// === OCP ===
	amount := 1000.0
	calculatorGold := DiscountCalculator{strategy: GoldCustomer{}}
	fmt.Println("Сумма со скидкой (Gold):", calculatorGold.Calculate(amount))

	// === ISP ===
	human := HumanWorker{}
	human.Work()
	human.Eat()
	human.Sleep()

	robot := RobotWorker{}
	robot.Work()

	// === DIP ===
	email := Notification{sender: EmailService{}}
	email.SendNotification("Добро пожаловать!")

	sms := Notification{sender: SmsService{}}
	sms.SendNotification("Ваш код: 123456")
}
