package task15

import "fmt"

func Run() {
	a := 1
	b := 2
	fmt.Println("a = ", a, "; b = ", b)

	a, b = b, a
	fmt.Println("a = ", a, "; b = ", b)
}
