Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
от 1 до 8 потом бесконечно 0

```
Программа выводит числа от 1 до 8, а затем продолжает выводить нули потому, что канал c, 
в который записываются значения из каналов a и b, не закрывается.
В функции merge используется бесконечный цикл for, 
который постоянно прослушивает каналы a и b и записывает значение в канал c, 
как только оно становится доступным. Однако, когда каналы a и b закрываются, 
блок select в функции merge все равно будет выполняться и ожидать чтения из закрытых каналов, 
что в итоге приводит к получению бесконечного потока нулей.