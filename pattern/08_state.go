package pattern

import "fmt"

// Интерфейс состояния
type state interface {
	handle() state
}

// Структура контекста
type context struct {
	state state
}

// Метод изменения состояния
func (c *context) changeState(s state) {
	c.state = s
}

// Реализация состояния StateA
type stateA struct{}

func (s *stateA) handle() state {
	fmt.Println("State A")
	return &stateB{}
}

// Реализация состояния StateB
type stateB struct{}

func (s *stateB) handle() state {
	fmt.Println("State B")
	return &stateC{}
}

// Реализация состояния StateC
type stateC struct{}

func (s *stateC) handle() state {
	fmt.Println("State C")
	return &stateA{}
}

func StateBuild() {
	// Создаем новый контекст с начальным состоянием StateA
	c := &context{state: &stateA{}}

	// Обрабатываем состояния в цикле
	for i := 0; i < 10; i++ {
		c.state = c.state.handle()
	}
}

/*
Этот код реализует состояния StateA, StateB и StateC,
где каждое состояние определено в отдельной структуре
и реализует интерфейс state с единственным методом handle().
Контекст содержит переменную state, которая указывает на текущее состояние.
В методе changeState() контекст может изменить текущее состояние на другое переданное состояние.
В функции StateBuild() создается новый контекст с начальным состоянием StateA,
а затем в цикле состояния обрабатываются с помощью метода handle().
В этом примере состояния просто выводят на экран свои имена, но на практике они могут выполнять любую другую логику.
*/
