package task26

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Run() {
	test := "this is a test string. Это тестовая строка. Абоба."

	fmt.Println("Before reverse:", test)
	test = reverse(test)
	fmt.Println("After reverse:", test)
}
