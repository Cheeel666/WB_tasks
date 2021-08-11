package task4

import (
	"fmt"
	"time"
)

func worker(ch <-chan interface{}, num int) {
	for i := range ch {
		fmt.Println("Worker ", num+1, ": ", i)
	}
}

// func loadData(ch chan<- interface{}) {
// 	response, err := http.Get("https://golangbyexample.com/iterate-over-a-string-golang/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	html, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, rn := range html {
// 		ch <- string(rn)
// 	}
// }

func timer(ch chan<- int) {
	time.Sleep(time.Second * 1)
	ch <- 0
}

func Run() {
	var n int
	fmt.Println("Введите количество рабочих:")
	fmt.Scanf("%d", &n)

	// Пока делал задание узнал про замечательную вещь - интерфейсный тип данных, решил попробовать. Всем советую!
	channel := make(chan interface{})
	endChannel := make(chan int)

	// Тут была попытка распарсить любой сайт по каналам, но потом я дочитал, что по условию, надо чтобы они еще
	// и закрывались когда нужно, поэтому я решил просто закинуть таймер в другой поток
	// go loadData(channel)

	// Тут собственно и находится таймер и рабочие
	go timer(endChannel)
	for i := 0; i < n; i++ {
		go worker(channel, i)
	}

	for {
		select {
		case channel <- "First case":
		case channel <- "Second case":
		case channel <- "Third case":
		// Когда мы получим данные из канала эндченнел, работа программы завершится.
		case <-endChannel:
			return
		}
	}

}
