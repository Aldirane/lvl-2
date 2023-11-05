package main

import "fmt"

// Определяем интерфейс для продукта
type Product interface {
	Use()
}

// Реализуем конкретные продукты
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

// Определяем интерфейс для фабрики
type Creator interface {
	CreateProduct() Product
}

// Реализуем конкретные фабрики
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Используем фабрики для создания продуктов
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	productA.Use()

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	productB.Use()
}

/*
В этом примере определены интерфейсы Product и Creator.
Конкретные продукты ConcreteProductA и ConcreteProductB реализуют интерфейс Product,
а конкретные фабрики ConcreteCreatorA и ConcreteCreatorB реализуют интерфейс Creator.
Метод CreateProduct фабрик возвращают конкретные продукты. Затем созданные продукты могут быть использованы через метод Use().
В функции main используются фабрики для создания продуктов Product и их последующего использования.
Надеюсь, это поможет вам понять, как реализовать паттерн Фабричный метод на Go.
*/
