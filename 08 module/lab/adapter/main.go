package main

import (
	"fmt"
)

type IPaymentProcessor interface {
	ProcessPayment(amount float64)
	RefundPayment(amount float64)
}

type InternalPaymentProcessor struct{}

func (p *InternalPaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Обработка платежа на сумму %.2f через внутреннюю систему.\n", amount)
}

func (p *InternalPaymentProcessor) RefundPayment(amount float64) {
	fmt.Printf("Возврат платежа на сумму %.2f через внутреннюю систему.\n", amount)
}

type ExternalPaymentSystemA struct{}

func (s *ExternalPaymentSystemA) MakePayment(amount float64) {
	fmt.Printf("Совершение платежа на %.2f через Внешнюю Систему А.\n", amount)
}

func (s *ExternalPaymentSystemA) MakeRefund(amount float64) {
	fmt.Printf("Совершение возврата на %.2f через Внешнюю Систему А.\n", amount)
}

type ExternalPaymentSystemB struct{}

func (s *ExternalPaymentSystemB) SendPayment(amount float64) {
	fmt.Printf("Отправка платежа на %.2f через Внешнюю Систему Б.\n", amount)
}

func (s *ExternalPaymentSystemB) ProcessRefund(amount float64) {
	fmt.Printf("Обработка возврата на %.2f через Внешнюю Систему Б.\n", amount)
}

//Адаптеры

type PaymentAdapterA struct {
	systemA *ExternalPaymentSystemA 
}

func NewPaymentAdapterA(systemA *ExternalPaymentSystemA) *PaymentAdapterA {
	return &PaymentAdapterA{systemA: systemA}
}

func (a *PaymentAdapterA) ProcessPayment(amount float64) {
	a.systemA.MakePayment(amount)
}

func (a *PaymentAdapterA) RefundPayment(amount float64) {
	a.systemA.MakeRefund(amount)
}

type PaymentAdapterB struct {
	systemB *ExternalPaymentSystemB 
}

func NewPaymentAdapterB(systemB *ExternalPaymentSystemB) *PaymentAdapterB {
	return &PaymentAdapterB{systemB: systemB}
}

func (a *PaymentAdapterB) ProcessPayment(amount float64) {
	a.systemB.SendPayment(amount)
}

func (a *PaymentAdapterB) RefundPayment(amount float64) {
	a.systemB.ProcessRefund(amount)
}

func GetPaymentProcessor(region string) IPaymentProcessor {
	switch region {
	case "RU":
		fmt.Println("Логика выбора: Выбрана внутренняя платежная система (Регион: RU).")
		return &InternalPaymentProcessor{}
	case "US":
		fmt.Println("Логика выбора: Выбрана Внешняя Система А (Регион: US).")
		return NewPaymentAdapterA(&ExternalPaymentSystemA{})
	case "EU":
		fmt.Println("Логика выбора: Выбрана Внешняя Система Б (Регион: EU).")
		return NewPaymentAdapterB(&ExternalPaymentSystemB{})
	default:
		fmt.Println("Логика выбора: Регион не опознан, используется внутренняя система (по умолчанию).")
		return &InternalPaymentProcessor{}
	}
}

func main() {
	fmt.Println("\nЗадание 2: Паттерн Адаптер ---")
	fmt.Println()

	fmt.Println("--- Сценарий 1: Регион RU ---")
	processorRU := GetPaymentProcessor("RU")
	processorRU.ProcessPayment(1500.0)
	processorRU.RefundPayment(200.0)
	fmt.Println("--------------------")

	fmt.Println("--- Сценарий 2: Регион US ---")
	processorUS := GetPaymentProcessor("US")
	processorUS.ProcessPayment(250.75)
	processorUS.RefundPayment(50.15)
	fmt.Println("--------------------")

	fmt.Println("--- Сценарий 3: Регион EU ---")
	processorEU := GetPaymentProcessor("EU")
	processorEU.ProcessPayment(999.0)
	processorEU.RefundPayment(150.0)
	fmt.Println("--------------------")

	fmt.Println("--- Сценарий 4: Регион JP (неизвестный) ---")
	processorDefault := GetPaymentProcessor("JP")
	processorDefault.ProcessPayment(5000.0)
	fmt.Println("--------------------")
}