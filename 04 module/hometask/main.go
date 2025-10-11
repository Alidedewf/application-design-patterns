package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IVehicle interface {
	Drive()
	Refuel()
}

type Car struct {
	Brand    string
	Model    string
	FuelType string
}

func (c Car) Drive() {
	fmt.Printf("Едем на автомобиле %s %s\n", c.Brand, c.Model)
}
func (c Car) Refuel() {
	fmt.Printf("Заправка автомобиля %s топливом: %s\n", c.Model, c.FuelType)
}

type Motorcycle struct {
	Type   string
	Engine int
}

func (m Motorcycle) Drive() {
	fmt.Printf("Едем на %s мотоцикле с двигателем %dcc\n", m.Type, m.Engine)
}
func (m Motorcycle) Refuel() {
	fmt.Println("Заправка мотоцикла бензином")
}

type Truck struct {
	Capacity   int
	AxlesCount int
}

func (t Truck) Drive() {
	fmt.Printf("Грузовик едет с грузом до %d кг\n", t.Capacity)
}
func (t Truck) Refuel() {
	fmt.Println("Заправка грузовика дизельным топливом")
}

type Bus struct {
	Seats int
	Route string
}

func (b Bus) Drive() {
	fmt.Printf("Автобус едет по маршруту %s с %d местами\n", b.Route, b.Seats)
}
func (b Bus) Refuel() {
	fmt.Println("Заправка автобуса дизельным топливом")
}

type VehicleFactory interface {
	CreateVehicle() IVehicle
}

type CarFactory struct {
	Brand    string
	Model    string
	FuelType string
}

func (f CarFactory) CreateVehicle() IVehicle {
	return Car{Brand: f.Brand, Model: f.Model, FuelType: f.FuelType}
}

type MotorcycleFactory struct {
	Type   string
	Engine int
}

func (f MotorcycleFactory) CreateVehicle() IVehicle {
	return Motorcycle{Type: f.Type, Engine: f.Engine}
}

type TruckFactory struct {
	Capacity   int
	AxlesCount int
}

func (f TruckFactory) CreateVehicle() IVehicle {
	return Truck{Capacity: f.Capacity, AxlesCount: f.AxlesCount}
}

type BusFactory struct {
	Seats int
	Route string
}

func (f BusFactory) CreateVehicle() IVehicle {
	return Bus{Seats: f.Seats, Route: f.Route}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Выберите тип транспорта (car, motorcycle, truck, bus): ")
	vehicleType, _ := reader.ReadString('\n')
	vehicleType = strings.TrimSpace(vehicleType)

	var factory VehicleFactory

	switch vehicleType {
	case "car":
		fmt.Print("Введите марку автомобиля: ")
		brand, _ := reader.ReadString('\n')
		fmt.Print("Введите модель автомобиля: ")
		model, _ := reader.ReadString('\n')
		fmt.Print("Введите тип топлива: ")
		fuel, _ := reader.ReadString('\n')

		factory = CarFactory{
			Brand:    strings.TrimSpace(brand),
			Model:    strings.TrimSpace(model),
			FuelType: strings.TrimSpace(fuel),
		}

	case "motorcycle":
		fmt.Print("Введите тип мотоцикла (спортивный, туристический): ")
		motoType, _ := reader.ReadString('\n')
		fmt.Print("Введите объем двигателя (в cc): ")
		engineStr, _ := reader.ReadString('\n')
		engine, _ := strconv.Atoi(strings.TrimSpace(engineStr))

		factory = MotorcycleFactory{
			Type:   strings.TrimSpace(motoType),
			Engine: engine,
		}

	case "truck":
		fmt.Print("Введите грузоподъемность (в кг): ")
		capStr, _ := reader.ReadString('\n')
		capacity, _ := strconv.Atoi(strings.TrimSpace(capStr))

		fmt.Print("Введите количество осей: ")
		axlesStr, _ := reader.ReadString('\n')
		axles, _ := strconv.Atoi(strings.TrimSpace(axlesStr))

		factory = TruckFactory{
			Capacity:   capacity,
			AxlesCount: axles,
		}

	case "bus":
		fmt.Print("Введите количество мест: ")
		seatsStr, _ := reader.ReadString('\n')
		seats, _ := strconv.Atoi(strings.TrimSpace(seatsStr))

		fmt.Print("Введите маршрут: ")
		route, _ := reader.ReadString('\n')

		factory = BusFactory{
			Seats: seats,
			Route: strings.TrimSpace(route),
		}

	default:
		fmt.Println("Неизвестный тип транспорта.")
		return
	}

	vehicle := factory.CreateVehicle()
	vehicle.Drive()
	vehicle.Refuel()
}