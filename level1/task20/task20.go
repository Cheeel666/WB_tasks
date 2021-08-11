package task20

import "fmt"

func Run() {
	slice := []string{"a", "a"}

	func(slice []string) {
		// Тут создается новая локальная переменная slice
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println("Inside func:", slice)
	}(slice)

	// тут slice = []string{"a", "a"}, т.к. глобальная область видимости
	fmt.Println("Outside func:", slice)
}
