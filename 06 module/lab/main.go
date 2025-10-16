package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IShippingStrategy interface {
	CalculateShippingCost(weight, distance float64) float64
	Name() string
}

type StandardShippingStrategy struct{}

func (s StandardShippingStrategy) CalculateShippingCost(weight, distance float64) float64 {
	return weight*0.5 + distance*0.1
}

func (s StandardShippingStrategy) Name() string {
	return "Standard"
}

type ExpressShippingStrategy struct{}

func (e ExpressShippingStrategy) CalculateShippingCost(weight, distance float64) float64 {
	return (weight*0.5 + distance*0.2) + 10
}

func (e ExpressShippingStrategy) Name() string {
	return "Express"
}

type InternationalShippingStrategy struct {
}

func (i InternationalShippingStrategy) CalculateShippingCost(weight, distance float64) float64 {
	return weight*1.0 + distance*0.5 + 15
}

func (i InternationalShippingStrategy) Name() string {
	return "International"
}

type NightShippingStrategy struct {
	Base     IShippingStrategy
	NightFee float64
}

func (n NightShippingStrategy) CalculateShippingCost(weight, distance float64) float64 {
	base := 0.0
	if n.Base != nil {
		base = n.Base.CalculateShippingCost(weight, distance)
	}
	return base + n.NightFee
}

func (n NightShippingStrategy) Name() string {
	return "Night"
}

type DeliveryContext struct {
	strategy IShippingStrategy
}

func (c *DeliveryContext) SetShippingStrategy(strategy IShippingStrategy) {
	c.strategy = strategy
}

func (c *DeliveryContext) CalculateCost(weight, distance float64) (float64, error) {
	if c.strategy == nil {
		return 0, errors.New("стратегия доставки не установлена")
	}
	if weight < 0 {
		return 0, errors.New("вес не может быть отрицательным")
	}
	if distance < 0 {
		return 0, errors.New("расстояние не может быть отрицательным")
	}
	return c.strategy.CalculateShippingCost(weight, distance), nil
}

func readLine(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func parseFloatInput(input string) (float64, error) {
	f, err := strconv.ParseFloat(strings.ReplaceAll(input, ",", "."), 64)
	if err != nil {
		return 0, errors.New("некорректное число")
	}
	return f, nil
}

func main() {
	fmt.Println("Программа расчета стоимости доставки")

	runSamples()

	deliveryContext := &DeliveryContext{}

	fmt.Println("\nВыберите тип доставки: 1 - Стандартная, 2 - Экспресс, 3 - Международная, 4 - Ночная")
	choice, err := readLine("Ваш выбор: ")
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	switch choice {
	case "1":
		deliveryContext.SetShippingStrategy(StandardShippingStrategy{})
	case "2":
		deliveryContext.SetShippingStrategy(ExpressShippingStrategy{})
	case "3":
		deliveryContext.SetShippingStrategy(InternationalShippingStrategy{})
	case "4":
		deliveryContext.SetShippingStrategy(NightShippingStrategy{Base: StandardShippingStrategy{}, NightFee: 5.0})
	default:
		fmt.Println("Неверный выбор.")
		return
	}

	wStr, err := readLine("Введите вес посылки (кг): ")
	if err != nil {
		fmt.Println("Ошибка чтения веса:", err)
		return
	}
	weight, err := parseFloatInput(wStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	dStr, err := readLine("Введите расстояние доставки (км): ")
	if err != nil {
		fmt.Println("Ошибка чтения расстояния:", err)
		return
	}
	distance, err := parseFloatInput(dStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	cost, err := deliveryContext.CalculateCost(weight, distance)
	if err != nil {
		fmt.Println("Ошибка расчета:", err)
		return
	}

	fmt.Printf("Стратегия: %s\n", deliveryContext.strategy.Name())
	fmt.Printf("Стоимость доставки: %.2f\n", cost)
}

func runSamples() {
	fmt.Println("\nДемонстрация работы стратегий")
	samples := []struct {
		str      IShippingStrategy
		weight   float64
		distance float64
	}{
		{StandardShippingStrategy{}, 2.0, 100.0},
		{ExpressShippingStrategy{}, 2.0, 100.0},
		{InternationalShippingStrategy{}, 2.0, 100.0},
		{NightShippingStrategy{Base: StandardShippingStrategy{}, NightFee: 5.0}, 2.0, 100.0},
		{StandardShippingStrategy{}, 5.5, 10.0},
		{ExpressShippingStrategy{}, 0.5, 3.0},
	}

	for _, s := range samples {
		c := s.str.CalculateShippingCost(s.weight, s.distance)
		fmt.Printf("%s: вес=%.2f кг, расстояние=%.2f км -> стоимость=%.2f\n", s.str.Name(), s.weight, s.distance, c)
	}
}
