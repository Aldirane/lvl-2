package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов
в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select,
которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции,
которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“done after %v”, time.Since(start))
*/

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			c <- 1
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(1*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		wg := new(sync.WaitGroup)
		mu := new(sync.RWMutex)
		wg.Add(len(channels))
		countGo := 0
		for i, ch := range channels {
			fmt.Printf("Go # %d\n", i+1)
			go func(ch <-chan interface{}, mu *sync.RWMutex, countGo *int, i int) {
				for val := range ch {
					fmt.Printf("Goroutine # %d done\n", i+1)
					done <- val
					mu.Lock()
					*countGo += 1
					mu.Unlock()
				}
			}(ch, mu, &countGo, i)
		}
		go func(wg *sync.WaitGroup, mu *sync.RWMutex, countGo *int) {
			for {
				mu.RLock()
				if *countGo > 0 {
					fmt.Printf("count = %d\n", *countGo)
					for i := 0; i < len(channels); i++ {
						wg.Done()
					}
					break
				}
				mu.RUnlock()
				time.Sleep(1 * time.Second)
			}
		}(wg, mu, &countGo)
		wg.Wait()
	}()
	return done
}
