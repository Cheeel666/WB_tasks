package task33

import (
	"fmt"
	"math/rand"
	"time"
)

func randNum(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		ch <- rand.Intn(100)
	}
}

func isEven(rand <-chan int, evenNums chan<- int) {
	for v := range rand {
		if v%2 == 0 {
			evenNums <- v
		}
	}
}

func output(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
func Run() {
	rand := make(chan int)
	evenNums := make(chan int)

	go randNum(rand)
	go isEven(rand, evenNums)
	go output(evenNums)
	time.Sleep(time.Millisecond * 500)
}
