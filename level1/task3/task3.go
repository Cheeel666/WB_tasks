package task3

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

func sumVal(arr chan int, res chan<- int) {
	sum := 0
	for i := range arr {
		sum += i
	}
	res <- sum
	close(res)
}

// Run runs task3
func Run() {
	fmt.Println("Задание: Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(22+32+42….) с использованием конкурентных вычислений.")
	fmt.Println()

	arr := [5]int{2, 4, 6, 8, 10}
	// Тут все также как и во втором задании, за исключением третьего канала, который также работает конкурентно
	arrVal := make(chan int)
	squares := make(chan int)
	res := make(chan int)

	go loadVal(arr, arrVal)
	go sqareVal(arrVal, squares)
	go sumVal(squares, res)

	for i := range res {
		fmt.Println(i)
	}
	fmt.Println()

}
