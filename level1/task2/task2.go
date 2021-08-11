package task2

import "fmt"

func loadVal(arr [5]int, c chan<- int) {
	for _, i := range arr {
		c <- i
		// fmt.Println("load")
	}

	close(c)
}

func sqareVal(arr chan int, res chan<- int) {
	for i := range arr {
		res <- i * i
		// fmt.Println("square")
	}

	close(res)
}

func Run() {
	fmt.Println("Задание: Написать программу, которая конкурентно рассчитает значение квадратов значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.")
	fmt.Println()

	arr := [5]int{2, 4, 6, 8, 10}

	arrVal := make(chan int)
	res := make(chan int)

	// Будем конкурентно загружать данные в канал и выводить их из него и возводить
	//  в квадрат, загружая в другой канал
	go loadVal(arr, arrVal)
	go sqareVal(arrVal, res)

	for i := range res {
		fmt.Println(i)
	}
	fmt.Println()
}
