Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

```
В данном случае строка "error" будет выведена потому, 
что переменная типа error не сравнивается с nil напрямую, а вызывается метод Error() на этой переменной.
В методе test() возвращается значение nil типа *customError. 
В переменную err, которая имеет тип error, записывается возвращаемое значение из test(). 
Поскольку *customError удовлетворяет интерфейсу error, 
то происходит неявное преобразование типа *customError к типу error. 
Здесь теряется информация о конкретном типе переменной и остается только интерфейсный тип error.
Когда происходит проверка if err != nil, по факту вызывается неявно метод Error() 
на переменной err (если err != nil), чтобы получить строковое представление ошибки. 
Таким образом, пустая переменная типа *customError = nil при вызове метода Error() возвращает пустую строку (""), 
и условие if err != nil выполняется, выводя строку "error".