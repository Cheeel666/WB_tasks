package theory

import (
	"fmt"
)

// дебаг показал, что вложенная функция, как и в anotherTest выполняется после ретюрна,
// но, я предполагаю, что из - за области видимости здесь значение нужной переменной инкрементируется,
// а в anotherTest инкреминтируется значение локальной переменной
func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func Task2() {
	fmt.Println(test())
	fmt.Println("AnotherTest:", anotherTest())
}
