/*
Visitor - это поведенческий паттерн проектирования,
который позволяет добавлять новые операции к существующей структуре объектов, не изменяя эти объекты.
Реализация паттерна Visitor состоит из следующих основных элементов:

	Интерфейс Visitor - определяет методы посещения для каждого типа элементов,
	которые могут быть посещены. Название метода может быть любым,
	но обычно используется один метод для каждого типа элемента.

	Интерфейс Element - определяет метод Accept,
	который принимает посетителя. Каждый тип элемента должен реализовывать этот метод.

Типы элементов - представляют собой различные объекты, которые могут быть посещены посетителем.
Каждый тип элемента должен иметь свою реализацию метода Accept, который передает себя посетителю.
Конкретный посетитель - реализует методы посещения для каждого типа элементов.
Эти методы обрабатывают каждый тип элемента по-разному.
*/

package main

import "fmt"

// Интерфейс Visitor
type Visitor interface {
	VisitConcreteElementA(element ConcreteElementA)
	VisitConcreteElementB(element ConcreteElementB)
}

// Интерфейс Element
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ConcreteElementA struct{}

func (e ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// Конкретный элемент B
type ConcreteElementB struct{}

func (e ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v ConcreteVisitor) VisitConcreteElementA(element ConcreteElementA) {
	fmt.Println("Посетитель посещает элемент A")
}

func (v ConcreteVisitor) VisitConcreteElementB(element ConcreteElementB) {
	fmt.Println("Посетитель посещает элемент B")
}

func main() {
	visitor := ConcreteVisitor{}

	elementA := ConcreteElementA{}
	elementB := ConcreteElementB{}

	// Вызываем метод Accept для каждого элемента
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}

/*
В данном примере мы создаем интерфейс Visitor, интерфейс Element и их конкретные реализации.
Затем создается конкретный посетитель ConcreteVisitor,
который реализует методы посещения для каждого типа элементов (ConcreteElementA и ConcreteElementB).

В функции main создается экземпляр посетителя и элементов,
а затем вызывается метод Accept для каждого элемента, передавая в него посетителя.

При выполнении программы будет выведено:
	Посетитель посещает элемент A
	Посетитель посещает элемент B
Таким образом, паттерн Visitor позволяет отделить операции, выполняемые над элементами,
от самих элементов, делая добавление новых операций более гибким и обеспечивая открытость для расширения.
*/
