package main

import "fmt"

// Coffee represents the base component interface for different types of coffee.
type Coffee interface {
	Cost() float64
	Description() string
}

// Espresso is a concrete component representing a specific type of coffee.
type Espresso struct{}

func (e *Espresso) Cost() float64 {
	return 1.5
}

func (e *Espresso) Description() string {
	return "Espresso"
}

// CoffeeDecorator is the base decorator interface.
type CoffeeDecorator interface {
	Coffee
}

// MilkDecorator is a concrete decorator that adds milk to the coffee.
type MilkDecorator struct {
	Coffee Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.Coffee.Cost() + 0.5
}

func (m *MilkDecorator) Description() string {
	return m.Coffee.Description() + ", Milk"
}

// SugarDecorator is a concrete decorator that adds sugar to the coffee.
type SugarDecorator struct {
	Coffee Coffee
}

func (s *SugarDecorator) Cost() float64 {
	return s.Coffee.Cost() + 0.25
}

func (s *SugarDecorator) Description() string {
	return s.Coffee.Description() + ", Sugar"
}

func main() {
	// Ordering an Espresso without any decorations.
	espresso := &Espresso{}
	fmt.Println("Base coffee:", espresso.Description())
	fmt.Println("Cost:", espresso.Cost())

	// Ordering an Espresso with Milk.
	espressoWithMilk := &MilkDecorator{Coffee: espresso}
	fmt.Println("Coffee with Milk:", espressoWithMilk.Description())
	fmt.Println("Cost:", espressoWithMilk.Cost())

	// Ordering an Espresso with Milk and Sugar.
	espressoWithMilkAndSugar := &SugarDecorator{Coffee: espressoWithMilk}
	fmt.Println("Coffee with Milk and Sugar:", espressoWithMilkAndSugar.Description())
	fmt.Println("Cost:", espressoWithMilkAndSugar.Cost())
}
