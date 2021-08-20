package theory

import (
	"fmt"
)

func Task1() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]

	fmt.Println("b:", b, len(b), cap(b))
	b[1] = 1
	fmt.Println("a:", a, len(a), cap(a)) // b ссылается на a, при смене значений в b меняются и в a
	b = append(b, 3, 3)
	fmt.Println("Увеличили capacity массива b в два раза - до 8")
	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("b:", b, len(b), cap(b))
	fmt.Println("Меняем первый элемент массива b")
	b[0] = 1
	fmt.Println("a:", a, len(a), cap(a)) //В массиве a все осталось по прежнему,
	//	так как изменилось capacity и b ссылается на другой массив
	fmt.Println("b:", b, len(b), cap(b))

}
