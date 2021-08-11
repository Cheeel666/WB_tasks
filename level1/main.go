package main

import (
	"fmt"
	"wb_homework/task1"
	"wb_homework/task10"
	"wb_homework/task11"
	"wb_homework/task12"
	"wb_homework/task13"
	"wb_homework/task14"
	"wb_homework/task15"
	"wb_homework/task16"
	"wb_homework/task17"
	"wb_homework/task18"
	"wb_homework/task19"
	"wb_homework/task2"
	"wb_homework/task20"
	"wb_homework/task21"
	"wb_homework/task22"
	"wb_homework/task23"
	"wb_homework/task24"
	"wb_homework/task25"
	"wb_homework/task26"
	"wb_homework/task27"
	"wb_homework/task28"
	"wb_homework/task29"
	"wb_homework/task3"
	"wb_homework/task30"
	"wb_homework/task31"
	"wb_homework/task32"
	"wb_homework/task33"
	"wb_homework/task34"
	"wb_homework/task4"
	"wb_homework/task5"
	"wb_homework/task7"
	"wb_homework/task8"
	"wb_homework/task9"
)

func run(num int) {
	switch num {
	case 1:
		task1.Run()
	case 2:
		task2.Run()
	case 3:
		task3.Run()
	case 4:
		task4.Run()
	case 5:
		task5.Run()
	case 6:
		fmt.Println("TODO")
	case 7:
		task7.Run()
	case 8:
		task8.Run()
	case 9:
		task9.Run()
	case 10:
		task10.Run()
	case 11:
		task11.Run()
	case 12:
		task12.Run()
	case 13:
		task13.Run()
	case 14:
		task14.Run()
	case 15:
		task15.Run()
	case 16:
		task16.Run()
	case 17:
		task17.Run()
	case 18:
		task18.Run()
	case 19:
		task19.Run()
	case 20:
		task20.Run()
	case 21:
		task21.Run()
	case 22:
		task22.Run()
	case 23:
		task23.Run()
	case 24:
		task24.Run()
	case 25:
		task25.Run()
	case 26:
		task26.Run()
	case 27:
		task27.Run()
	case 28:
		task28.Run()
	case 29:
		task29.Run()
	case 30:
		task30.Run()
	case 31:
		task31.Run()
	case 32:
		task32.Run()
	case 33:
		task33.Run()
	case 34:
		task34.Run()
	}
}

func main() {
	var i int
	fmt.Println("Input number:")
	_, err := fmt.Scanf("%d", &i)

	for {
		run(i)
		fmt.Println("Input number:")
		_, err = fmt.Scanf("%d", &i)
		if err != nil {
			break
		}
	}
}
