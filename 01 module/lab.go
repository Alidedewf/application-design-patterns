package main

import "fmt"

type Vehicle struct {
	ID            string
	brand         string
	model         string
	year          int
	enginestarted bool
}

type Car struct {
	Vehicle
	count        int
	transmission string
}

type Motorcycle struct {
	Vehicle
	body string
	box  string
}

type Garage struct {
	ID       string
	Vehicles []Vehicle
}

type Fleet struct {
	Garages []Garage
}

func (v *Vehicle) StartEngine() {
	if v.enginestarted {
		fmt.Println("Двигатель уже запущен")
	} else {
		v.enginestarted = true
		fmt.Printf("Двигатель %s %s запущен\n", v.brand, v.model)
	}
}

func (g *Garage) AddVehicle(v Vehicle) {
	g.Vehicles = append(g.Vehicles, v)
}

func (f *Fleet) AddGarage(g Garage) {
	f.Garages = append(f.Garages, g)
}

func (f *Fleet) RemoveGarage(garageID string) {
	for i, g := range f.Garages {
		if g.ID == garageID {
			f.Garages = append(f.Garages[:i], f.Garages[i+1:]...)
			return
		}

	}
}

func (f *Fleet) FindVehicle(vehicleID string) (*Vehicle, *Garage) {
	for i, garage := range f.Garages {
		for j, v := range garage.Vehicles {
			if v.ID == vehicleID {
				return &f.Garages[i].Vehicles[j], &f.Garages[i]
			}
		}
	}
	return nil, nil
}

func main() {
	v1 := Vehicle{ID: "01", brand: "Toyota", model: "Camry", year: 2018}
	v2 := Vehicle{ID: "02", brand: "Hyndai", model: "Accent", year: 2019}
	v3 := Vehicle{ID: "03", brand: "Geely", model: "Emgrand", year: 2025}
	v4 := Vehicle{ID: "04", brand: "Hyndai", model: "Accent", year: 2014}

	garage01 := Garage{ID: "01"}
	garage02 := Garage{ID: "02"}

	garage01.AddVehicle(v1)
	garage01.AddVehicle(v2)
	garage02.AddVehicle(v3)
	garage02.AddVehicle(v4)

	fleet := Fleet{}
	fleet.AddGarage(garage01)
	fleet.AddGarage(garage02)

	Vehicle, garage := fleet.FindVehicle("02")
	if Vehicle != nil {
		fmt.Printf("Найдено транспортное средство: %s %s в гараже под номером %s\n", Vehicle.brand, Vehicle.model, garage.ID)
	} else {
		fmt.Printf("Транспортное средство не найдено")
	}
	fleet.RemoveGarage("02")
}
