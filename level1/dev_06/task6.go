package task6

import (
	"fmt"
	"time"
)

// 1) С помощью дополнительного канал
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

// 2) Передача переменной по ссылке
func f(flag *bool) {
	for {
		if *flag {
			fmt.Println("working")
		} else {
			break
		}
	}
	fmt.Println("stop")
}

func Run() {
	mainChannel := make(chan int)
	waitChannel := make(chan bool)
	go writer(mainChannel, waitChannel)
	go waiter(waitChannel, 1)

	// // flag := true
	// // go f(&flag)

	time.Sleep(time.Second)

	// // flag = false

	//Также можно просто закрыть канал, но тогда прога упадет
}
