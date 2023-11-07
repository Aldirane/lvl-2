/*
Паттерн строитель (builder) позволяет конструировать сложные объекты шаг за шагом,
с использованием одного и того же конструирующего кода.
*/

package pattern

import "fmt"

// Структура объекта, который мы хотим построить
type Product struct {
	partA, partB, partC string
}

// Интерфейс строителя
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() Product
}

// Конкретный строитель
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "Part B"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "Part C"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

// Директор - отвечает за шаги построения объекта
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func BuilderBuild() {
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}

	// Строим объект с помощью директора
	director.Construct()
	product := builder.GetResult()

	// Выводим результат
	fmt.Printf("Part A: %s\n", product.partA)
	fmt.Printf("Part B: %s\n", product.partB)
	fmt.Printf("Part C: %s\n", product.partC)
}

/*
Этот пример демонстрирует паттерн строитель с использованием интерфейсов в Go.
Создается структура Product, которая представляет объект, который мы хотим построить.
Затем создается интерфейс Builder, определяющий основные методы для построения объекта.
В данном примере реализуется конкретный строитель ConcreteBuilder,
который имплементирует интерфейс Builder и строит объект типа Product.
Директор Director отвечает за последовательность шагов построения объекта,
а именно вызов методов BuildPartA(), BuildPartB() и BuildPartC() у строителя.
В функции BuilderBuild() создается экземпляр строителя ConcreteBuilder,
который затем передается в директор и через директор строится объект типа Product.
Результат можно увидеть в выводе.
*/
