package main

import (
	"fmt"
	"strings"
)

// 1 задача:
// Написать программу переводящую строку в int число без использование стандартных библиотек

// strRev reverces string
func strRev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func myPow(a, b int) int {
	var res int
	res = 1
	for b > 0 {
		res *= a
		b--
	}
	return res
}

func sToInt(str string) int {
	strRune := []rune(strRev(str))
	var res int = 0
	for i := 0; i < len(strRune); i++ {
		// fmt.Println(int(strRune[i]))

		res += (int(strRune[i]) - 48) * (myPow(10, i))
	}
	return res
}

// 2 задача:
// Написать программу находящую палиндром строковой и числовой
func polindrom(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		// fmt.Println(runes[i], runes[len(runes)-i-1])
		if runes[i] == runes[len(runes)-i-1] {

			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	var str string
	// var num int
	str = "1488"
	fmt.Println("Задача 1:")
	fmt.Println(sToInt(str))

	fmt.Println("Задача 2:")
	strs := []string{"Я иду с мечем судия", "12344321", "22", "132354", "ksgjshjgjsg", "ss"}
	for _, s := range strs {
		fmt.Println(polindrom(s))
	}

}
