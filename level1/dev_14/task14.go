package task14

import "fmt"

func makeSet(s [5]string) map[string]bool {
	var resultSet = make(map[string]bool)
	for i := 0; i < len(s); i++ {
		resultSet[s[i]] = true
	}
	return resultSet
}

func Run() {
	str := [5]string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Последовательность:", str)
	fmt.Println("Множество:", makeSet(str))
}
