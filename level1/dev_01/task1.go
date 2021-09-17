package task1

import "fmt"

// Human implements human
type Human struct {
	name string
	age  int
}

// Action implements human action
type Action struct {
	Human
	madeAction bool
}

// NewAction implements composition of structs
func NewAction(name string, age int, makeAction bool) *Action {
	return &Action{
		Human: Human{
			name: name,
			age:  age,
		},
		madeAction: makeAction,
	}
}

// Run implements run task function
func Run() {
	act := NewAction("ilya", 21, true)
	fmt.Println("Задание: Реализовать композицию структуры Action от родительской структуры Human.")
	fmt.Println()
	fmt.Println(*act)
	fmt.Println()
}
