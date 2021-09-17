package theory

import (
	"fmt"
)

func Task6() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s) // поменялось только первое число,
	// потому что внутри функции указатель на слайс стал указывать на другой слайс
}

func modifySlice(i []string) {
	i[0] = "3"
	fmt.Println("fist:", len(i), cap(i))

	i = append(i, "4") // i больше не ссылается на слайс i, который в свою очередь ссылается на массив,
	// так как capacity удвоилась
	// и i ссылается уже на другой массив(скопировав значения из первого массива),
	// отличный от от того на который ссылался первоначальный i

	fmt.Println("after append", len(i), cap(i))
	fmt.Println(i)

	i[1] = "5"
	fmt.Println("after change 1st value", len(i), cap(i))
	fmt.Println(i)

	i = append(i, "6")
	fmt.Println("after second append", len(i), cap(i))
	fmt.Println(i)
}
