package task1

import "fmt"

// Структура человека
type Human struct {
	name string
	age  int
}

// Структура действия, которая в качестве поля использует человека
type Action struct {
	Human
	madeAction bool
}

// NewAction - композиция структур
func NewAction(name string, age int, makeAction bool) *Action {
	return &Action{
		Human: Human{
			name: name,
			age:  age,
		},
		madeAction: makeAction,
	}
}

func Run() {
	act := NewAction("ilya", 21, true)
	fmt.Println("Задание: Реализовать композицию структуры Action от родительской структуры Human.")
	fmt.Println()
	fmt.Println(*act)
	fmt.Println()
}
