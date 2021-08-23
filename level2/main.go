package main

// import (
// 	"level2/patterns"
// 	"level2/theory"
// import t "github.com/Cheeel666/WB_tasks/level2/practice/task1"

// )

import (
	"errors"
	"fmt"
	"strconv"
)

// multChar multiples char
func multChar(char string, n int) string {
	resultStr := ""
	for i := 0; i < n; i++ {
		resultStr += char
	}
	return resultStr
}

// UnpackStr unpacks string
func UnpackStr(str string) (string, error) {
	var result string
	var flag bool
	flag = false

	if len(str) == 0 {
		return "", nil
	}

	_, err := strconv.Atoi(str)
	if err == nil {
		return "", errors.New("Некорректная строка")
	}

	_, err = strconv.Atoi(string(str[0]))
	if err == nil {
		return "", errors.New("Некорректная строка")
	}

	strRunes := []rune(str)
	currentChar := strRunes[0]

	for i := 0; i < len(strRunes); i++ {
		n, err := strconv.Atoi(string(strRunes[i]))
		if rune(str[i]) == '\\' {
			flag = true
		}
		if err == nil || flag {
			result += multChar(string(currentChar), n-1)
			currentChar = strRunes[i+1]
			flag = false
		} else {
			result += string(strRunes[i])
			currentChar = strRunes[i]
		}
	}
	return result, nil
}
func main() {
	// patterns.Facade()
	// patterns.BuilderPattern()
	// patterns.VisitorPattern()
	// patterns.CommandPattern()
	// patterns.ChainOfResponsibilityPattern()
	// patterns.FactoryMethodPattern()
	// patterns.StrategyPattern()
	// patterns.StatePattern()

	// theory.Task1()
	// theory.Task2()
	// theory.Task3()
	// theory.Task4()
	// theory.Task5()
	// theory.Task6()
	// theory.Task7()
	// t.TimeNow("")
	fmt.Println(UnpackStr("qwe\4\5"))
}
