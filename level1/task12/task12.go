package task12

import "fmt"

func update(p *int) {
	b := 2
	p = &b
	// *p = b // в данном случае p будет указывать на b
}

func Run() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) //Выводим значение, на которое указывает p. В данном случае 1
	update(p)
	fmt.Println(*p) // p по прежнему указывает на a
}
