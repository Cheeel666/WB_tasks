package task11

import "fmt"

func Run() {
	firstArr := [5]int{1, 3, 2, 0, 5}
	secondArr := [7]int{2, 1, 0, 4, 5, 6, 7}

	for _, i := range firstArr {
		for _, j := range secondArr {
			if i == j {
				fmt.Println(i)
			}
		}
	}
}
