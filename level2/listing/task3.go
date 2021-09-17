package theory

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil // fs.PathError - интерфейс, присваиваем значение nil
	return err
}

func Task3() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)         //пустой интерфейс != nil
	xType := fmt.Sprintf("%T", err) // Тип интерфейса
	nilType := fmt.Sprintf("%T", nil)
	fmt.Println(xType) // Ну тут видно уже
	fmt.Println(nilType)
}
