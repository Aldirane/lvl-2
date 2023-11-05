package main

import "fmt"

type Request struct {
	data string
}

type Handler interface {
	SetNext(handler Handler)
	Handle(request Request)
}

type AbstractHandler struct {
	nextHandler Handler
}

func (ah *AbstractHandler) SetNext(handler Handler) {
	ah.nextHandler = handler
}

type ConcreteHandler1 struct {
	AbstractHandler
}

func (ch1 *ConcreteHandler1) Handle(request Request) {
	if request.data == "request1" {
		fmt.Println("Handled by ConcreteHandler1")
		return
	}
	if ch1.nextHandler != nil {
		ch1.nextHandler.Handle(request)
	}
}

type ConcreteHandler2 struct {
	AbstractHandler
}

func (ch2 *ConcreteHandler2) Handle(request Request) {
	if request.data == "request2" {
		fmt.Println("Handled by ConcreteHandler2")
		return
	}
	if ch2.nextHandler != nil {
		ch2.nextHandler.Handle(request)
	}
}

type ConcreteHandler3 struct {
	AbstractHandler
}

func (ch3 *ConcreteHandler3) Handle(request Request) {
	if request.data == "request3" {
		fmt.Println("Handled by ConcreteHandler3")
		return
	}
	if ch3.nextHandler != nil {
		ch3.nextHandler.Handle(request)
	}
}

func main() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	request1 := Request{data: "request1"}
	handler1.Handle(request1)

	request2 := Request{data: "request2"}
	handler1.Handle(request2)

	request3 := Request{data: "request3"}
	handler1.Handle(request3)

	request4 := Request{data: "request4"}
	handler1.Handle(request4)
}

/*
В этом примере созданы несколько обработчиков (ConcreteHandler1, ConcreteHandler2 и ConcreteHandler3),
которые имеют ссылку на следующий обработчик через абстрактный обработчик (AbstractHandler).
При вызове метода Handle каждый обработчик проверяет, может ли он обработать запрос.
Если он может, то выводит сообщение, что он обработал запрос.
Если не может, то передает запрос следующему обработчику.
Если нет следующего обработчика, то запрос не будет обработан ни одним из обработчиков.
В main создаются обработчики и устанавливается последовательность их вызова (handler1 -> handler2 -> handler3).
Далее создаются четыре запроса и вызывается у первого обработчика метод Handle для каждого запроса.
*/
