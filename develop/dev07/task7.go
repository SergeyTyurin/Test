package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

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

fmt.Printf(“fone after %v”, time.Since(start))
*/

import (
	"fmt"
	"sync"
	"time"
)

func joinChannels(channels ...<-chan interface{}) <-chan interface{} {
	singleCh := make(chan interface{})
	if len(channels) == 0 {
		close(singleCh)
		return singleCh
	}
	go func() {
		var isClosed bool
		var mu sync.Mutex

		for _, ch := range channels {
			// проверка каждого канала, закрыт он или нет.
			go func(ch <-chan interface{}) {
				for {
					_, isChanOpened := <-ch
					if !isChanOpened {
						break
					}
				}

				// Синхронизируем проверку, что канал закрыт и закрытие канала для горутин, проверяющих закрытие входных каналов
				mu.Lock()
				if !isClosed {
					// закрываем объединенный канал, когда первый из каналов, поступивших на вход закроется
					close(singleCh)
				}
				isClosed = true
				mu.Unlock()
			}(ch)
		}

	}()

	return singleCh
}

func main() {
	// функция закрывает канал по истечении времени after
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	// объединяем несколько каналов в один. (ожидаем первый закрытый канал)
	singleChan := joinChannels(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	// Ожидаем, пока объединенный канал не закроется
	<-singleChan

	// печатаем информацию
	fmt.Printf("done after %v\n", time.Since(start))

}
