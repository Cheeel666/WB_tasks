package task27

import (
	"fmt"
	"strings"
)

func reverseStr(str string) string {
	words := strings.Split(str, " ")
	len := len(words)
	for i := 0; i < len/2; i++ {
		tmp := words[len-i-1]
		words[len-i-1] = words[i]
		words[i] = tmp
	}
	str = strings.Join(words, " ")
	return str
}
func Run() {
	str := "This is a test string"

	fmt.Println("String before reverse: ", str)
	str = reverseStr(str)
	fmt.Println("String after reverse: ", str)
}
