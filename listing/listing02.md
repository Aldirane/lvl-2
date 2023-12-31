Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```
defer выполняет вызов функции при возврате из функции в которой он используется.
Эти вызовы defer добавляются в стек приложения, которые
срабатывают в порядке от последнего к первому.
При выходе из функции test в return мы явно не возвращаем x, но так как
она была инициализирована в шапке возвращаемых типов значений функции test.
То ответ 2, так как x инкрементируется вызовом анонимной функции после return.
При выходе из anotherTest в return указана переменная x, 
которая не была инициализирована как в test,
поэтому defer инкрементирует локальную копию x, но не ту,
что была возвращена из функции.
