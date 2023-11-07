package pattern

import "fmt"

// Интерфейс стратегии
type SortStrategy interface {
	Sort([]int) []int
}

// Реализация конкретной стратегии - сортировка по возрастанию
type AscendingSortStrategy struct{}

func (s AscendingSortStrategy) Sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// Реализация конкретной стратегии - сортировка по убыванию
type DescendingSortStrategy struct{}

func (s DescendingSortStrategy) Sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// Контекст, использующий стратегию
type SortContext struct {
	strategy SortStrategy
}

func (c *SortContext) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

func (c *SortContext) Sort(arr []int) []int {
	return c.strategy.Sort(arr)
}

func StrategyBuild() {
	// Создание контекста с стратегией сортировки по возрастанию
	context := SortContext{}
	context.SetStrategy(AscendingSortStrategy{})
	arr := []int{5, 2, 8, 3, 1}
	sortedArr := context.Sort(arr)
	fmt.Println(sortedArr)

	// Изменение стратегии на сортировку по убыванию
	context.SetStrategy(DescendingSortStrategy{})
	sortedArr = context.Sort(arr)
	fmt.Println(sortedArr)
}

/*
В этом примере определены две конкретные стратегии сортировки: AscendingSortStrategy и DescendingSortStrategy,
реализующие метод Sort сортировки массива чисел по возрастанию и убыванию соответственно.
Контекст SortContext использует интерфейс SortStrategy, который позволяет установить стратегию сортировки и вызвать метод Sort.
В функции StrategyBuild создается контекст с начальной стратегией сортировки по возрастанию,
а затем вызывается метод Sort для сортировки массива.
Затем изменяется стратегия сортировки на DescendingSortStrategy и выполняется снова метод Sort.
Выходные данные будут:
[1, 2, 3, 5, 8]
[8, 5, 3, 2, 1]
Таким образом, применение паттерна Strategy позволяет динамически изменять поведение объекта
в зависимости от выбранной стратегии сортировки.
*/
