package main

import (
	"fmt"
)

type PaymentStrategy interface {
	Pay(amount float64)
}

type CardPayment struct{}

func (c *CardPayment) Pay(amount float64) {
	fmt.Printf("💳 Оплата %.2f₸ банковской картой прошла успешно.\n", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("💸 Оплата %.2f₸ через PayPal выполнена успешно.\n", amount)
}

type CryptoPayment struct{}

func (c *CryptoPayment) Pay(amount float64) {
	fmt.Printf("🪙 Оплата %.2f₸ криптовалютой подтверждена в блокчейне.\n", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) {
	if p.strategy == nil {
		fmt.Println("Ошибка: способ оплаты не выбран!")
		return
	}
	p.strategy.Pay(amount)
}

type Observer interface {
	Update(currency string, rate float64)
}

type Subject interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type CurrencyExchange struct {
	observers     []Observer
	currencyRates map[string]float64
}

func NewCurrencyExchange() *CurrencyExchange {
	return &CurrencyExchange{
		observers:     []Observer{},
		currencyRates: make(map[string]float64),
	}
}

func (c *CurrencyExchange) AddObserver(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *CurrencyExchange) RemoveObserver(o Observer) {
	for i, observer := range c.observers {
		if observer == o {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *CurrencyExchange) SetRate(currency string, rate float64) {
	c.currencyRates[currency] = rate
	fmt.Printf("\n📈 Новый курс %s: %.2f₸\n", currency, rate)
	c.NotifyObservers()
}

func (c *CurrencyExchange) NotifyObservers() {
	for _, observer := range c.observers {
		for currency, rate := range c.currencyRates {
			observer.Update(currency, rate)
		}
	}
}

type Bank struct{}

func (b *Bank) Update(currency string, rate float64) {
	fmt.Printf("🏦 Банк обновил курс для %s: %.2f₸\n", currency, rate)
}

type Investor struct{}

func (i *Investor) Update(currency string, rate float64) {
	fmt.Printf("💰 Инвестор получил уведомление: %s теперь %.2f₸\n", currency, rate)
}

type NewsAgency struct{}

func (n *NewsAgency) Update(currency string, rate float64) {
	fmt.Printf("📰 Новости: курс %s изменился до %.2f₸\n", currency, rate)
}

func main() {
	var choice int
	var amount float64
	context := &PaymentContext{}

	fmt.Println("Выберите способ оплаты:")
	fmt.Println("1 - Банковская карта")
	fmt.Println("2 - PayPal")
	fmt.Println("3 - Криптовалюта")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		context.SetStrategy(&CardPayment{})
	case 2:
		context.SetStrategy(&PayPalPayment{})
	case 3:
		context.SetStrategy(&CryptoPayment{})
	default:
		fmt.Println("Неверный выбор.")
		return
	}

	fmt.Print("Введите сумму оплаты: ")
	fmt.Scan(&amount)
	context.ExecutePayment(amount)

	exchange := NewCurrencyExchange()

	bank := &Bank{}
	investor := &Investor{}
	news := &NewsAgency{}

	exchange.AddObserver(bank)
	exchange.AddObserver(investor)
	exchange.AddObserver(news)

	exchange.SetRate("USD", 486.50)
	exchange.SetRate("EUR", 512.80)
}
