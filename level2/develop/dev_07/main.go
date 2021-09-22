package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	res := make(chan interface{})

	// Выгружаем все из канала
	closeChan := func(channel <-chan interface{}) {
		for val := range channel {
			res <- val
		}
		// Ждем указанное в горутине время и помечаем что сделали
		wg.Done()
	}
	// Добавляем вейтгруппы
	wg.Add(len(channels))
	for _, channel := range channels {
		// Запускаем параллельно каналы
		go closeChan(channel)
	}
	go func() {
		// Тут ожидаем
		wg.Wait()
		close(res)
	}()
	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Millisecond),
		sig(50*time.Millisecond),
		sig(100*time.Millisecond),
		sig(1*time.Millisecond),
		sig(1*time.Millisecond),
	)

	fmt.Printf("done after %v", time.Since(start))
}
