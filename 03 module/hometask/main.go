package main

import "fmt"

// === SRP ===

type Order struct {
	ProductName string
	Quantity    int
	Price       float64
}

type PriceCalculator struct{}

func (pc *PriceCalculator) CalculateTotal(order Order) float64 {
	return order.Price * float64(order.Quantity) * 0.9
}

type PaymentProcessor struct{}

func (pp *PaymentProcessor) Process(paymentDetails string) {
	fmt.Println("Платеж обрабатывается с помощью:", paymentDetails)
}

type EmailNotifier struct{}

func (en EmailNotifier) SendEmail(email string) {
	fmt.Println("Письмо с подтверждением отправлено на адрес:", email)
}

// === OCP ===

type Employee interface {
	CalculateSalary() float64
}

type PermanentEmployee struct {
	BaseSalary float64
}

func (p PermanentEmployee) CalculateSalary() float64 {
	return p.BaseSalary * 1.2
}

type ContractEmployee struct {
	BaseSalary float64
}

func (c ContractEmployee) CalculateSalary() float64 {
	return c.BaseSalary * 1.1
}

type Intern struct {
	BaseSalary float64
}

func (i Intern) CalculateSalary() float64 {
	return i.BaseSalary * 0.8
}

func PrintSalary(e Employee) {
	fmt.Println("Зарплата:", e.CalculateSalary())
}

// === ISP ===

type Printer interface {
	Print(content string)
}

type Scanner interface {
	Scan(content string)
}

type Fax interface {
	Fax(content string)
}

type AllInOnePrinter struct{}

func (p AllInOnePrinter) Print(content string) {
	fmt.Println("Печатает:", content)
}
func (p AllInOnePrinter) Scan(content string) {
	fmt.Println("Сканирование:", content)
}
func (p AllInOnePrinter) Fax(content string) {
	fmt.Println("Факс:", content)
}

type BasicPrinter struct{}

func (p BasicPrinter) Print(content string) {
	fmt.Println("Печатает:", content)
}

// === DIP ===

type Notifier interface {
	Send(message string)
}

type EmailSender struct{}

func (e EmailSender) Send(message string) {
	fmt.Println("Письмо отправлено:", message)
}

type SmsSender struct{}

func (s SmsSender) Send(message string) {
	fmt.Println("СМС отправлен:", message)
}

type NotificationService struct {
	notifiers []Notifier
}

func (n *NotificationService) SendNotification(message string) {
	for _, notifier := range n.notifiers {
		notifier.Send(message)
	}
}

func main() {
	// SRP
	order := Order{"Laptop", 2, 1000}
	calculator := PriceCalculator{}
	total := calculator.CalculateTotal(order)
	fmt.Println("Total price:", total)

	payment := PaymentProcessor{}
	payment.Process("Visa 1234")

	emailNotifier := EmailNotifier{}
	emailNotifier.SendEmail("ali@example.com")

	// OCP
	PrintSalary(PermanentEmployee{BaseSalary: 1000})
	PrintSalary(Intern{BaseSalary: 1000})
	PrintSalary(ContractEmployee{BaseSalary: 1000})

	// ISP
	basic := BasicPrinter{}
	basic.Print("Test page")

	allInOne := AllInOnePrinter{}
	allInOne.Print("Page A")
	allInOne.Scan("Doc B")
	allInOne.Fax("Doc C")

	// DIP
	email := EmailSender{}
	sms := SmsSender{}
	notificationService := NotificationService{
		notifiers: []Notifier{email, sms},
	}
	notificationService.SendNotification("Order shipped!")
}
