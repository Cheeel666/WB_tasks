package task17

import "fmt"

func typeOf(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	default:
		fmt.Println("unknown")
	}
}

func Run() {
	var (
		i  int
		s  string
		b  bool
		ch chan int
	)
	typeOf(i)
	typeOf(s)
	typeOf(b)
	typeOf(ch)
}
