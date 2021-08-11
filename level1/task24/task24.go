package task24

import "fmt"

func Run() {
	slice := make([]int, 10, 100)
	fmt.Println("Result: ", slice, "with len: ", len(slice), "and cap: ", cap(slice))
}
