package task9

import "fmt"

func read(ch chan<- int, arr [5]int) {
	for _, i := range arr {
		ch <- i
	}
	// Если не закроем тут канал, возникнет дедлок
	close(ch)
}

func square(ch <-chan int, res chan<- int) {
	for i := range ch {
		res <- 2 * i
	}
	// Если не закроем тут канал, возникнит дедлок
	close(res)
}

func printSquares(ch <-chan int, ret chan bool) {
	for i := range ch {
		fmt.Println(i)
	}

	ret <- true
}
func Run() {
	arrChan := make(chan int)
	squareChan := make(chan int)
	retChan := make(chan bool)

	arr := [5]int{1, 2, 3, 4, 5}
	go read(arrChan, arr)
	go square(arrChan, squareChan)
	go printSquares(squareChan, retChan)

	fmt.Println("Done:", <-retChan)
}
