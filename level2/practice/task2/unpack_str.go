package task2

import (
	"errors"
	"strconv"
)

// strRev reverces string
func strRev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

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

	strRunes := []rune(strRev(str))
	n := 1

	for i := 0; i < len(strRunes); i++ {
		currentN, err := strconv.Atoi(string(strRunes[i]))

		if err == nil {
			n = currentN
		} else {
			result += multChar(string(strRunes[i]), n)
			n = 1
		}
	}
	return strRev(result), nil
}
