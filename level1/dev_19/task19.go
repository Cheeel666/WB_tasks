package task19

import "fmt"

var justString string

func createHugeString(n int) string {
	var s string = ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

// Не обьявлено CreateHugeString
func someFunc() {
	v := createHugeString(1 << 10) //двоичный сдвиг 1 << 10 = (1 * 2 - 10 раз) = 1024
	justString = v[:100]
	fmt.Println("V:", v)
	fmt.Println("justString:", justString)
}

// Предполагаю, что проблема здесь в глобальной переменной. Проблемы безопасности кода
func Run() {
	someFunc()
	fmt.Println("main:", justString)
}
