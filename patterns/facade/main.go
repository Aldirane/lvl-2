/*
Паттерн фасад представляет собой структурный паттерн проектирования,
который позволяет предоставить простой интерфейс для более сложной системы структур.
*/

package main

import "fmt"

// Внутренние модули или подсистемы
type ModuleA struct{}

func (a *ModuleA) MethodA() {
	fmt.Println("Вызван метод MethodA модуля ModuleA")
}

type ModuleB struct{}

func (b *ModuleB) MethodB() {
	fmt.Println("Вызван метод MethodB модуля ModuleB")
}

type ModuleC struct{}

func (c *ModuleC) MethodC() {
	fmt.Println("Вызван метод MethodC модуля ModuleC")
}

// Фасад
type Facade struct {
	moduleA *ModuleA
	moduleB *ModuleB
	moduleC *ModuleC
}

func NewFacade() *Facade {
	return &Facade{
		moduleA: &ModuleA{},
		moduleB: &ModuleB{},
		moduleC: &ModuleC{},
	}
}

// Упрощенные методы фасада для доступа к сложным системам извне
func (f *Facade) Operation1() {
	fmt.Println("Operation1 запущена")
	f.moduleA.MethodA()
	f.moduleB.MethodB()
	fmt.Println("Operation1 завершена")
}

func (f *Facade) Operation2() {
	fmt.Println("Operation2 запущена")
	f.moduleB.MethodB()
	f.moduleC.MethodC()
	fmt.Println("Operation2 завершена")
}

// Использование фасада
func main() {
	facade := NewFacade()

	// Выполнение сложных операций с использованием фасада
	facade.Operation1()
	facade.Operation2()
}

/*
В данном примере у нас есть три модуля (ModuleA, ModuleB, ModuleC),
каждый из которых имеет свои собственные методы. Затем мы создаем фасад,
который объединяет эти модули и предоставляет упрощенный интерфейс для их использования.

Методы Operation1 и Operation2 являются упрощенными методами фасада,
которые вызывают методы соответствующих модулей.
При использовании фасада мы можем выполнять сложные операции с помощью этих методов,
скрывая сложность взаимодействия с отдельными модулями.

Результат выполнения программы будет следующим:

Operation1 запущена
Вызван метод MethodA модуля ModuleA
Вызван метод MethodB модуля ModuleB
Operation1 завершена
Operation2 запущена
Вызван метод MethodB модуля ModuleB
Вызван метод MethodC модуля ModuleC
Operation2 завершена
Видно, что мы можем выполнить сложную операцию (например, Operation1)
с помощью простого вызова метода фасада,
и он позаботится о вызове необходимых методов внутри себя,
скрывая сложность взаимодействия с подсистемой.
*/
