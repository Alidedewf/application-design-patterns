package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type beverage interface {
	brew() error
	addCondiments() error
	customerWantsCondiments() bool
}


func boilWater() error {
	fmt.Println("Кипячение воды...")
	return nil
}

func pourInCup() error {
	fmt.Println("Наливание в чашку...")
	return nil
}

func prepareRecipe(b beverage) error {
	fmt.Println("--- Начало приготовления ---")
	
	if err := boilWater(); err != nil {
		return fmt.Errorf("ошибка кипячения: %w", err)
	}

	if err := b.brew(); err != nil {
		return fmt.Errorf("ошибка заваривания: %w", err)
	}

	if err := pourInCup(); err != nil {
		return fmt.Errorf("ошибка наливания: %w", err)
	}

	if b.customerWantsCondiments() {
		if err := b.addCondiments(); err != nil {
			return fmt.Errorf("ошибка добавления ингредиентов: %w", err)
		}
	} else {
		fmt.Println("Ингредиенты пропущены по желанию клиента.")
	}

	fmt.Println("--- Напиток готов! ---")
	return nil
}

type Tea struct {
	HasLemon bool
}

func (t *Tea) brew() error {
	fmt.Println("Заваривание чая...")
	return nil
}

func (t *Tea) addCondiments() error {
	if !t.HasLemon {
		return fmt.Errorf("нет лимона для чая")
	}
	fmt.Println("Добавление лимона...")
	return nil
}

func (t *Tea) customerWantsCondiments() bool {
	return true
}

type Coffee struct {
	HasMilk bool
	MilkType string
}

func (c *Coffee) brew() error {
	fmt.Println("Заваривание кофе...")
	return nil
}

func (c *Coffee) addCondiments() error {
	if !c.HasMilk {
		return fmt.Errorf("нет молока для кофе")
	}
	fmt.Printf("Добавление сахара и %s молока...\n", c.MilkType)
	return nil
}

func (c *Coffee) customerWantsCondiments() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Хотите добавить сахар и молоко? (y/n): ")

	for {
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))
		
		if answer == "y" {
			return true
		}
		if answer == "n" {
			return false
		}
		fmt.Print("Пожалуйста, введите 'y' или 'n': ")
	}
}

type HotChocolate struct {
	HasMarshmallows bool
}

func (hc *HotChocolate) brew() error {
	fmt.Println("Растворение какао-порошка в молоке...")
	return nil
}

func (hc *HotChocolate) addCondiments() error {
	if !hc.HasMarshmallows {
		return fmt.Errorf("нет зефира")
	}
	fmt.Println("Добавление сливок...")
	return nil
}

func (hc *HotChocolate) customerWantsCondiments() bool {
	return true
}

func main() {
	fmt.Println("-- Приготовление чая --")
	tea := &Tea{HasLemon: true}
	if err := prepareRecipe(tea); err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	fmt.Println()

	fmt.Println("-- Приготовление кофе --")
	coffee := &Coffee{HasMilk: true, MilkType: "соевого"}
	if err := prepareRecipe(coffee); err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	fmt.Println()

	fmt.Println("-- Приготовление горячего шоколада --")
	hotChoc := &HotChocolate{HasMarshmallows: true}
	if err := prepareRecipe(hotChoc); err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	fmt.Println()

	fmt.Println("-- Отсуствие ингредиентов --")
	teaNoLemon := &Tea{HasLemon: false}
	if err := prepareRecipe(teaNoLemon); err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}
}