Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
```
При сравнении во втором Println вывод false потому что
у них разные типы данных. У переменной err тип *os.PathError
а у nil тип nil. Поэтому они не равны.
но если мы сделам так,
```go
errVal := fmt.Sprintf("%v", err)
nilVal := fmt.Sprintf("%v", nil)
fmt.Println(errVal == nilVal)
```
то будет выведено true потому что сравниваются значения а не типы.