package main

import "fmt"

type ITransport interface {
	Move()
	FuelUp()
}

type Car struct {
	Model string
}

func (c Car) Move() {
	fmt.Printf("Автомобиль %s движется по дороге\n", c.Model)
}

func (c Car) FuelUp() {
	fmt.Printf("Автомобиль %s заправляется бензином.\n", c.Model)
}

type Motorcycle struct {
	Model string
}

func (m Motorcycle) Move() {
	fmt.Printf("Мопед %s движется по дороге\n", m.Model)
}

func (m Motorcycle) FuelUp() {
	fmt.Printf("Мопед %s заправляется бензином \n", m.Model)
}

type Plane struct {
	Airline string
}

func (p Plane) Move() {
	fmt.Printf("Самолет из %s летит в небе.\n", p.Airline)
}

func (p Plane) FuelUp() {
	fmt.Printf("Самолет из %s заправляется топливом.\n", p.Airline)
}

type TransportFactory interface {
	CreateTransport() ITransport
}

type CarFactory struct{}
type MotorcycleFactory struct{}
type PlaneFactory struct{}

func (f CarFactory) CreateTransport() ITransport {
	return Car{Model: "Hyndai"}
}

func (f MotorcycleFactory) CreateTransport() ITransport {
	return Motorcycle{Model: "Suzuki"}
}

func (f PlaneFactory) CreateTransport() ITransport {
	return Plane{Airline: "AirAstana"}
}
