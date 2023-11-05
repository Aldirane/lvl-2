package main

import "fmt"

// Интерфейс команды
type Command interface {
	Execute()
}

// Конкретная команда
type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Получатель команды
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Выполнение команды")
}

// Инициатор команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) Run() {
	i.command.Execute()
}

func main() {
	// Создание экземпляров объектов
	receiver := Receiver{}
	command := &ConcreteCommand{receiver: receiver}
	invoker := Invoker{}

	// Настройка и запуск команды
	invoker.SetCommand(command)
	invoker.Run()
}

/*
В этом примере показана реализация паттерна команда с использованием интерфейса Command и конкретной команды ConcreteCommand.
Receiver представляет получателя команды, который выполняет фактическое действие, когда команда вызывается.
Invoker является инициатором команды, который устанавливает команду и запускает ее выполнение.
В функции main создаются экземпляры всех объектов и настраивается команда с помощью Invoker.
Затем команда выполняется, вызывая метод Run() у Invoker, который запускает выполнение команды.
В результате на консоль будет выведено сообщение "Выполнение команды".
*/
