package task4

import (
	"sort"
	"strings"
)

// makeElem makes an element for the result map
func makeElem(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

// stringInSlice checks if word in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// sortMap forms resulting map
func sortMap(testMap map[string][]string) map[string][]string {
	resultMap := make(map[string][]string)
	for _, w := range testMap {
		resultMap[w[0]] = w
		sort.Strings(w)
	}
	return resultMap
}

// MakeSets makes sets with anagrams
func MakeSets(str []string) map[string][]string {
	testMap := make(map[string][]string)
	var checkSlice []string

	for i := 0; i < len(str); i++ {
		word := strings.ToLower(str[i])
		letters := makeElem(word)
		if stringInSlice(letters, checkSlice) {
			testMap[letters] = append(testMap[letters], str[i])
		} else {
			checkSlice = append(checkSlice, letters)
			testMap[letters] = []string{str[i]}
		}
	}

	testMap = sortMap(testMap)
	return testMap
}
