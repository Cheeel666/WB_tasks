package main

import (
	"fmt"
	"wb_homework/dev_01"
	"wb_homework/dev_02"
	"wb_homework/dev_03"
	"wb_homework/dev_04"
	"wb_homework/dev_05"
	"wb_homework/dev_06"
	"wb_homework/dev_07"
	"wb_homework/dev_08"
	"wb_homework/dev_09"
	"wb_homework/dev_10"
	"wb_homework/dev_11"
	"wb_homework/dev_12"
	"wb_homework/dev_13"
	"wb_homework/dev_14"
	"wb_homework/dev_15"
	"wb_homework/dev_16"
	"wb_homework/dev_17"
	"wb_homework/dev_18"
	"wb_homework/dev_19"
	"wb_homework/dev_20"
	"wb_homework/dev_21"
	"wb_homework/dev_22"
	"wb_homework/dev_23"
	"wb_homework/dev_24"
	"wb_homework/dev_25"
	"wb_homework/dev_26"
	"wb_homework/dev_27"
	"wb_homework/dev_28"
	"wb_homework/dev_29"
	"wb_homework/dev_30"
	"wb_homework/dev_31"
	"wb_homework/dev_32"
	"wb_homework/dev_33"
	"wb_homework/dev_34"
)

func run(num int) {
	switch num {
	case 1:
		dev_01.Run()
	case 2:
		dev_02.Run()
	case 3:
		dev_03.Run()
	case 4:
		dev_04.Run()
	case 5:
		dev_05.Run()
	case 6:
		dev_06.Run()
	case 7:
		dev_07.Run()
	case 8:
		dev_08.Run()
	case 9:
		dev_09.Run()
	case 10:
		dev_10.Run()
	case 11:
		dev_11.Run()
	case 12:
		dev_12.Run()
	case 13:
		dev_13.Run()
	case 14:
		dev_14.Run()
	case 15:
		dev_15.Run()
	case 16:
		dev_16.Run()
	case 17:
		dev_17.Run()
	case 18:
		dev_18.Run()
	case 19:
		dev_19.Run()
	case 20:
		dev_20.Run()
	case 21:
		dev_21.Run()
	case 22:
		dev_22.Run()
	case 23:
		dev_23.Run()
	case 24:
		dev_24.Run()
	case 25:
		dev_25.Run()
	case 26:
		dev_26.Run()
	case 27:
		dev_27.Run()
	case 28:
		dev_28.Run()
	case 29:
		dev_29.Run()
	case 30:
		dev_30.Run()
	case 31:
		dev_31.Run()
	case 32:
		dev_32.Run()
	case 33:
		dev_33.Run()
	case 34:
		dev_34.Run()
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
