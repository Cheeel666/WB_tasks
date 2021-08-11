package task16

import "fmt"

func Run() {
	n := 0
	if true {
		n := 1 // Здесь мы создали локальную переменную которая видна в пределах {}
		n++
	}
	fmt.Println(n)
}
