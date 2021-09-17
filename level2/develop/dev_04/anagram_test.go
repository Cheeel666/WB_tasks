package task4

import (
	"fmt"
	"testing"
)

func TestMakeSets(t *testing.T) {
	var str [3][]string
	str[0] = []string{"абоба", "амогус"}
	str[1] = []string{"австриец", "ганец", "царствие", "агнец"}
	str[2] = []string{"аймак", "кайма", "майка", "анкилоз", "кинозал", "козлина", "лозинка"}

	for i := 0; i < len(str); i++ {
		fmt.Println(MakeSets(str[i]))
	}
}
