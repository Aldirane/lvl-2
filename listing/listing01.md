Что выведет программа?
```go
package main

import (
    "fmt"
)

func main() {
    a:= [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```
Ответ
```go
[77 78 79]
```
Вывод индексов от 1 до 4 невключительно.