package task5

import (
	"fmt"
	"time"
)

func reader(ch <-chan int) {
	for i := range ch {
		fmt.Println("Readed from channel:", i)
	}
}

func writer(ch chan<- int, end <-chan bool) {
	for i := 0; ; i++ {
		select {
		case <-end:
			return
		default:
			ch <- i
		}
	}
}

func waiter(ch chan<- bool, n int) {
	// fmt.Println("strat")
	time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println("end")
	ch <- true
}

func Run() {
	var n int
	fmt.Println("Введите количество секунд:")
	fmt.Scanf("%d", &n)
	mainChannel := make(chan int)
	waitChannel := make(chan bool)

	go reader(mainChannel)
	go writer(mainChannel, waitChannel)
	go waiter(waitChannel, n)
	time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println(<-waitChannel)
	return
	// // самый простой способ, но тогда программа валится с паникой
	// time.Sleep(time.Duration(n) * time.Second)
	// close(mainChannel)
}
