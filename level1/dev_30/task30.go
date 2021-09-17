package task30

import "fmt"

func Run() {
	testslice := []int{1, 2, 3, 4, 5}

	i := 3

	fmt.Println("Test slice before cut of ", i, " element:", testslice)
	testslice = append(testslice[:i], testslice[i+1:]...)
	fmt.Println("Test slice after cut of ", i, " element:", testslice)
}
