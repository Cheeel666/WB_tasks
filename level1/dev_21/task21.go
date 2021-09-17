package task21

import (
	"fmt"
	"time"
)

func writeArray(arr [5]int, ch chan<- int) {
	for _, i := range arr {
		ch <- i
	}
}

func Run() {
	arr := [5]int{1, 2, 3, 4, 5}

	ch := make(chan int)

	go writeArray(arr, ch)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)
	time.Sleep(time.Second)
}
