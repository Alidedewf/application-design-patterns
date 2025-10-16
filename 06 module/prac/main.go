package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ICostCalculationStrategy interface {
	CalculateBaseCost(params BookingParams) (float64, error)
	Name() string
}

type ClassType string

const (
	Economy  ClassType = "economy"
	Business ClassType = "business"
)

type PassengerType string

const (
	Adult  PassengerType = "adult"
	Child  PassengerType = "child"
	Senior PassengerType = "senior"
)

type BookingParams struct {
	DistanceKm          float64
	Passengers          int
	Class               ClassType
	PassengerProfiles   []PassengerProfile
	BaggagePerPassenger int
	Transfers           int
	RegionMultiplier    float64
	ExtraServicesFee    float64
}

type PassengerProfile struct {
	Type        PassengerType
	DiscountPct float64
}

type PlaneStrategy struct{}

func (p PlaneStrategy) Name() string { return "Plane" }

func (p PlaneStrategy) CalculateBaseCost(params BookingParams) (float64, error) {
	if params.DistanceKm < 0 {
		return 0, errors.New("distance can't be negative")
	}
	base := 50.0 + params.DistanceKm*0.12
	fuelSurcharge := base * 0.08
	airportFees := 20.0 + float64(params.Transfers)*5.0
	return base + fuelSurcharge + airportFees, nil
}

type TrainStrategy struct{}

func (t TrainStrategy) Name() string { return "Train" }

func (t TrainStrategy) CalculateBaseCost(params BookingParams) (float64, error) {
	if params.DistanceKm < 0 {
		return 0, errors.New("distance can't be negative")
	}
	base := 10.0 + params.DistanceKm*0.05
	transferFee := float64(params.Transfers) * 2.0
	return base + transferFee, nil
}

type BusStrategy struct{}

func (b BusStrategy) Name() string { return "Bus" }

func (b BusStrategy) CalculateBaseCost(params BookingParams) (float64, error) {
	if params.DistanceKm < 0 {
		return 0, errors.New("distance can't be negative")
	}
	base := 5.0 + params.DistanceKm*0.03
	transferFee := float64(params.Transfers) * 1.0
	return base + transferFee, nil
}

type TravelBookingContext struct {
	strategy ICostCalculationStrategy
}

func NewTravelBookingContext() *TravelBookingContext {
	return &TravelBookingContext{strategy: nil}
}

func (c *TravelBookingContext) SetStrategy(s ICostCalculationStrategy) {
	c.strategy = s
}

func (c *TravelBookingContext) CalculateTotalCost(params BookingParams) (float64, error) {
	if c.strategy == nil {
		return 0, errors.New("no pricing strategy set")
	}
	if params.Passengers <= 0 {
		return 0, errors.New("passengers must be > 0")
	}
	if params.DistanceKm < 0 {
		return 0, errors.New("distance can't be negative")
	}

	basePerPassenger, err := c.strategy.CalculateBaseCost(params)
	if err != nil {
		return 0, err
	}

	classMultiplier := classMultiplier(params.Class)

	baggageFeePerUnit := 10.0

	total := 0.0
	for i := 0; i < params.Passengers; i++ {
		profile := PassengerProfile{Type: Adult, DiscountPct: 0}
		if i < len(params.PassengerProfiles) {
			profile = params.PassengerProfiles[i]
		}

		cost := basePerPassenger * classMultiplier

		cost += float64(params.BaggagePerPassenger) * baggageFeePerUnit

		if profile.DiscountPct > 0 {
			cost = cost * (1 - profile.DiscountPct/100.0)
		}

		switch profile.Type {
		case Child:
			cost *= 0.5
		case Senior:
			cost *= 0.8
		}

		total += cost
	}

	if params.Passengers >= 5 {
		total *= 0.9
	}

	total = total*params.RegionMultiplier + params.ExtraServicesFee

	return math.Round(total*100) / 100, nil
}

func classMultiplier(c ClassType) float64 {
	switch c {
	case Business:
		return 2.0
	default:
		return 1.0
	}
}

func main() {
	fmt.Println("=== Travel Booking Simulator (Strategy Pattern) ===")

	runSamples()

	ctx := NewTravelBookingContext()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Выберите транспорт: 1 - Самолёт, 2 - Поезд, 3 - Автобус")
	choice, _ := readLine(reader, "Ваш выбор: ")

	switch choice {
	case "1":
		ctx.SetStrategy(PlaneStrategy{})
	case "2":
		ctx.SetStrategy(TrainStrategy{})
	case "3":
		ctx.SetStrategy(BusStrategy{})
	default:
		fmt.Println("Неверный выбор. Выходим.")
		return
	}

	// Ввод параметров
	distanceStr, _ := readLine(reader, "Введите расстояние (км): ")
	distance, err := parsePositiveFloat(distanceStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	passStr, _ := readLine(reader, "Количество пассажиров: ")
	passengers, err := parsePositiveInt(passStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	classStr, _ := readLine(reader, "Класс (economy/business): ")
	classStr = strings.ToLower(strings.TrimSpace(classStr))
	var class ClassType
	if classStr == "business" {
		class = Business
	} else {
		class = Economy
	}

	baggageStr, _ := readLine(reader, "Багаж на пассажира (шт): ")
	baggage, err := parseNonNegativeInt(baggageStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	transfersStr, _ := readLine(reader, "Количество пересадок: ")
	transfers, err := parseNonNegativeInt(transfersStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	regionStr, _ := readLine(reader, "Региональный коэффициент (пример 1.0): ")
	region, err := parsePositiveFloat(regionStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	extraStr, _ := readLine(reader, "Доп. услуги (фикс. сумма): ")
	extra, err := parseNonNegativeFloat(extraStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Хотите ввести профили пассажиров для скидок? (y/n)")
	yn, _ := readLine(reader, "")
	profiles := []PassengerProfile{}
	if strings.ToLower(strings.TrimSpace(yn)) == "y" {
		for i := 0; i < passengers; i++ {
			prompt := fmt.Sprintf("Пассажир %d (adult/child/senior), персон. скидка %% (например 0): ", i+1)
			line, _ := readLine(reader, prompt)
			parts := strings.Fields(line)
			p := PassengerProfile{Type: Adult, DiscountPct: 0}
			if len(parts) >= 1 {
				t := strings.ToLower(parts[0])
				if t == "child" {
					p.Type = Child
				}
				if t == "senior" {
					p.Type = Senior
				}
			}
			if len(parts) >= 2 {
				d, err := strconv.ParseFloat(strings.ReplaceAll(parts[1], ",", "."), 64)
				if err == nil && d >= 0 {
					p.DiscountPct = d
				}
			}
			profiles = append(profiles, p)
		}
	}

	params := BookingParams{
		DistanceKm:          distance,
		Passengers:          passengers,
		Class:               class,
		PassengerProfiles:   profiles,
		BaggagePerPassenger: baggage,
		Transfers:           transfers,
		RegionMultiplier:    region,
		ExtraServicesFee:    extra,
	}

	cost, err := ctx.CalculateTotalCost(params)
	if err != nil {
		fmt.Println("Ошибка расчёта:", err)
		return
	}

	fmt.Printf("Стратегия: %s\n", ctx.strategy.Name())
	fmt.Printf("Общая стоимость для %d пассажиров: %.2f\n", passengers, cost)
}

func readLine(r *bufio.Reader, prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	text, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func parsePositiveFloat(s string) (float64, error) {
	s = strings.TrimSpace(strings.ReplaceAll(s, ",", "."))
	if s == "" {
		return 0, errors.New("пустой ввод")
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil || v <= 0 {
		return 0, errors.New("ожидалось положительное число")
	}
	return v, nil
}

func parseNonNegativeFloat(s string) (float64, error) {
	s = strings.TrimSpace(strings.ReplaceAll(s, ",", "."))
	if s == "" {
		return 0, nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil || v < 0 {
		return 0, errors.New("ожидалось неотрицательное число")
	}
	return v, nil
}

func parsePositiveInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 {
		return 0, errors.New("ожидалось положительное целое число")
	}
	return v, nil
}

func parseNonNegativeInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	v, err := strconv.Atoi(s)
	if err != nil || v < 0 {
		return 0, errors.New("ожидалось неотрицательное целое число")
	}
	return v, nil
}

func runSamples() {
	fmt.Println("\n-- Примеры расчётов --")
	ctx := NewTravelBookingContext()

	examples := []struct {
		strategy ICostCalculationStrategy
		params   BookingParams
		label    string
	}{
		{PlaneStrategy{}, BookingParams{DistanceKm: 1000, Passengers: 1, Class: Economy, PassengerProfiles: []PassengerProfile{{Type: Adult}}, BaggagePerPassenger: 1, Transfers: 1, RegionMultiplier: 1.1, ExtraServicesFee: 0}, "Plane, 1000km, economy, 1 passenger"},
		{PlaneStrategy{}, BookingParams{DistanceKm: 1000, Passengers: 2, Class: Business, PassengerProfiles: []PassengerProfile{{Type: Adult}, {Type: Senior, DiscountPct: 5}}, BaggagePerPassenger: 2, Transfers: 0, RegionMultiplier: 1.2, ExtraServicesFee: 50}, "Plane, business, 2 passengers (senior)"},
		{TrainStrategy{}, BookingParams{DistanceKm: 300, Passengers: 3, Class: Economy, PassengerProfiles: []PassengerProfile{{Type: Adult}, {Type: Child}, {Type: Adult}}, BaggagePerPassenger: 1, Transfers: 0, RegionMultiplier: 1.0, ExtraServicesFee: 0}, "Train, 300km, includes child"},
		{BusStrategy{}, BookingParams{DistanceKm: 50, Passengers: 6, Class: Economy, PassengerProfiles: []PassengerProfile{{Type: Adult}, {Type: Adult}, {Type: Adult}, {Type: Adult}, {Type: Adult}, {Type: Adult}}, BaggagePerPassenger: 0, Transfers: 1, RegionMultiplier: 1.0, ExtraServicesFee: 0}, "Bus, group of 6 (group discount)"},
	}

	for _, ex := range examples {
		ctx.SetStrategy(ex.strategy)
		cost, err := ctx.CalculateTotalCost(ex.params)
		if err != nil {
			fmt.Printf("%s -> Ошибка: %v\n", ex.label, err)
			continue
		}
		fmt.Printf("%s -> Total: %.2f (strategy=%s)\n", ex.label, cost, ctx.strategy.Name())
	}
}
