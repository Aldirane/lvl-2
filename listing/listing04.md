Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1                                                                  
2                                                                  
3                                                                  
4                                                                  
5                                                                  
6                                                                  
7                                                                  
8                                                                  
9                                                                  
fatal error: all goroutines are asleep - deadlock!                 
                                                                   
goroutine 1 [chan receive]:                                        
main.main()                                                        
        /home/aldar/my_projects/go_l2/test/linters/main.go:11 +0xa8
exit status 2                           

```
Функция main ожидает получение в канал ch, 
но после завершения всех горутин в цикле, уходит в бесконечное
ожидание так как некому отправить новые значения в канал ch.
Поэтому deadlock