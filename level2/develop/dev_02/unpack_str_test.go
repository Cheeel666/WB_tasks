package task2

import (
	"fmt"
	"testing"
)

func TestUnpackStr(t *testing.T) {
	var str [7]string
	str[0] = "a4bc2d5e"
	str[1] = "abcd"
	str[2] = "45"
	str[3] = ""
	str[4] = "4a5"
	str[5] = "a5A"
	str[6] = "a5"

	for i := 0; i < len(str); i++ {
		res, err := UnpackStr(str[i])
		if err == nil {
			fmt.Println("Test ", i+1, ": OK; result: ", res)
		} else {
			fmt.Println("Test ", i+1, "error:", err)
		}
	}
}
