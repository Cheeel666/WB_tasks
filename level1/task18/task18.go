package task18

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100 // Меняем значение 0 индекса
	v = append(v, b)
	fmt.Println("Slice inside function", v) // Из принта видно, что создается новый участок памяти,
	// на который ссылается слайс
}

func Run() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
