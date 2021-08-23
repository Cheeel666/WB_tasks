package task2

import (
	"errors"
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

		if err == nil {
			result += multChar(string(currentChar), n-1)
			currentChar = strRunes[i+1]
		} else {
			result += string(strRunes[i])
			currentChar = strRunes[i]
		}
	}
	return result, nil
}
