package theory

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test5() *customError {
	{
		// do something
	}
	return nil
}

func Task5() {
	var err error
	err = test5()
	fmt.Println(err)
	typeof := fmt.Sprintf("%T", err)
	fmt.Println(typeof) //*theory.customError
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
