package main

import "fmt"

type IBadWorker interface {
	Work()
	Eat()
	Sleep()
}

type Human struct{}

func (h *Human) Work() {
	fmt.Println("Человек работает")
}

func (h *Human) Eat() {
	fmt.Println("Человек ест")
}

func (h *Human) Sleep() {
	fmt.Println("Человек спит")
}

type Robot struct{}

func (r *Robot) Work() {
	fmt.Println("Робот работает")
}

func (r *Robot) Eat() {
	panic("Роботы не едят")
}

func (r *Robot) Sleep() {
	panic("Роботы не спят")
}

type Worker interface {
	Work()
}

type WorkerAdapter struct {
	worker IBadWorker
}

func NewWorkerAdapter(w IBadWorker) *WorkerAdapter {
	return &WorkerAdapter{worker: w}
}

func (a *WorkerAdapter) Work() {
	a.worker.Work()
}

func main() {
	human := &Human{}
	robot := &Robot{}

	var workers = []Worker{
		NewWorkerAdapter(human),
		NewWorkerAdapter(robot),
	}

	for _, w := range workers {
		w.Work()
	}
}