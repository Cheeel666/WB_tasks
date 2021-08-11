package task34

import "fmt"

// Просто кидаем в множество интовые значения символов из строки и сравниваем длину со строкой
func isUnique(s string) bool {
	var controlSet = make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		controlSet[s[i]] = true
	}
	// fmt.Println(len(s), ":", len(controlSet))
	return len(s) == len(controlSet)
}

func Run() {
	test := "aabcdsdear"
	testOne := "abcdefghijklmnop"
	fmt.Println("Test 1 string is unique:", isUnique(test))
	fmt.Println("Test 2 string is unique:", isUnique(testOne))

}
