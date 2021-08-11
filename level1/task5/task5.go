package task5

import (
	"fmt"
	"time"
)

func reader(ch <-chan int) {
	for _ = range ch {
		fmt.Println("Readed from channel:", <-ch)
	}
}

func writer(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// func waiter(ch chan<- int, n int) {
// 	time.Sleep(time.Duration(n) * time.Second)
// 	ch <- 1
// }

func Run() {
	var n int
	fmt.Println("Введите количество секунд:")
	fmt.Scanf("%d", &n)
	mainChannel := make(chan int)
	// waitChannel := make(chan int)

	go reader(mainChannel)
	go writer(mainChannel)
	// go waiter(waitChannel, n)

	// самый простой способ, но тогда программа валится с паникой
	time.Sleep(time.Duration(n) * time.Second)
	close(mainChannel)
}
