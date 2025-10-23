package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BeverageHook interface {
	Brew()
	AddCondiments()
	CustomerWantsCondiments() bool 
}

func boilWater() {
	fmt.Println("Кипятим воду")
}

func pourInCup() {
	fmt.Println("Наливаем в чашку")
}

func PrepareRecipe(b BeverageHook) {
	boilWater()
	
	b.Brew()
	
	pourInCup()
	
	if b.CustomerWantsCondiments() {
		b.AddCondiments()
	}
}

type Tea struct{}

func (t *Tea) Brew() {
	fmt.Println("Завариваем чай")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Добавляем лимон")
}

func (t *Tea) CustomerWantsCondiments() bool {
	return true
}

type Coffee struct{}

func (c *Coffee) Brew() {
	fmt.Println("Пропускаем кофе через фильтр")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Добавляем сахар и молоко")
}

func (c *Coffee) CustomerWantsCondiments() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Хотите добавить сахар и молоко? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "y" {
			return true
		}
		if answer == "n" {
			return false
		}
		fmt.Println("Некорректный ввод. Пожалуйста, введите (y/n).")
	}
}

type HotChocolate struct{}

func (hc *HotChocolate) Brew() {
	fmt.Println("Растворяем шоколадный порошок")
}

func (hc *HotChocolate) AddCondiments() {
	fmt.Println("Добавляем взбитые сливки")
}

func (hc *HotChocolate) CustomerWantsCondiments() bool {
	return true
}

func main() {
	fmt.Println("--- Готовим чай ---")
	tea := &Tea{}
	PrepareRecipe(tea)

	fmt.Println("\n--- Готовим кофе ---")
	coffee := &Coffee{}
	PrepareRecipe(coffee) 

	fmt.Println("\n--- Готовим какао---")
	hotChoc := &HotChocolate{}
	PrepareRecipe(hotChoc)
}