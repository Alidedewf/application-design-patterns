package main

import "fmt"

type IPaymentProcessor interface {
	ProcessPayment(amount float64)
}

type PayPalPaymentProcessor struct{}

func (p PayPalPaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("PayPal processed payment of $%.2f\n", amount)
}

type StripePaymentService struct{}

func (s StripePaymentService) MakeTransaction(totalAmount float64) {
	fmt.Printf("Stripe processed transaction of $%.2f\n", totalAmount)
}

type StripePaymentAdapter struct {
	stripe StripePaymentService
}

func (a StripePaymentAdapter) ProcessPayment(amount float64) {
	a.stripe.MakeTransaction(amount)
}

type CryptoPayService struct{}

func (c CryptoPayService) SendCrypto(amount float64) {
	fmt.Printf("Crypto payment of $%.2f sent\n", amount)
}

type CryptoPayAdapter struct {
	crypto CryptoPayService
}

func (a CryptoPayAdapter) ProcessPayment(amount float64) {
	a.crypto.SendCrypto(amount)
}

func testAdapter() {
	var payment IPaymentProcessor

	payment = PayPalPaymentProcessor{}
	payment.ProcessPayment(20.0)

	payment = StripePaymentAdapter{stripe: StripePaymentService{}}
	payment.ProcessPayment(45.5)

	payment = CryptoPayAdapter{crypto: CryptoPayService{}}
	payment.ProcessPayment(100.0)
}


func main() {
	fmt.Println("\n=== Adapter Pattern Demo ===")
	testAdapter()
}