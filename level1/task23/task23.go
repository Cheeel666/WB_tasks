package task23

import (
	"fmt"
	"sort"
)

func Run() {
	a := []string{"A", "D", "B", "C", "E"}
	x := "C"

	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено %s по индексу %d в %v.\n", x, i, a)
	} else {
		fmt.Printf("Не найдено %s в %v.\n", x, a)
	}
}
