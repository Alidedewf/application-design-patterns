package main

import "fmt"

type Employee struct {
	name string
	ID   string
	post string
}

type Worker struct {
	Employee
	hourly int
	hours  int
}

func (w *Worker) CalculateSalary() int {
	return w.hourly * w.hours
}

type Manager struct {
	Employee
	fix   int
	bonus int
}

func (m *Manager) CalculateSalary() int {
	return m.fix + m.bonus
}

func main() {
	worker1 := Worker{
		Employee: Employee{name: "Алишер", ID: "01", post: "Рабочий"},
		hourly:   2500,
		hours:    160,
	}
	worker2 := Worker{
		Employee: Employee{name: "Адилет", ID: "02", post: "Рабочий"},
		hourly:   2500,
		hours:    200,
	}

	manager1 := Manager{
		Employee: Employee{name: "Андрей", ID: "01", post: "Менеджер"},
		fix:      400000,
		bonus:    50000,
	}

	manager2 := Manager{
		Employee: Employee{name: "Иван", ID: "02", post: "Менеджер"},
		fix:      350000,
		bonus:    50000,
	}

	workers := []Worker{worker1, worker2}
	managers := []Manager{manager1, manager2}

	for _, w := range workers {
		fmt.Printf("ID: %s | %s (%s) Зарплата: %d\n", w.ID, w.name, w.post, w.CalculateSalary())
	}

	for _, m := range managers {
		fmt.Printf("ID: %s | %s (%s) Зарплата: %d\n", m.ID, m.name, m.post, m.CalculateSalary())
	}
}
