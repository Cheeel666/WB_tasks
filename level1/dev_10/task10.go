package task10

import "fmt"

func groupUp(nums ...float32) map[int][]float32 {
	mp := make(map[int][]float32)

	for _, num := range nums {
		mp[int(num/10)*10] = append(mp[int(num/10)*10], num)
	}
	return mp
}

func Run() {
	res := groupUp(32.5, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, -5, 5)

	fmt.Println("Result:", res)
	// fmt.Println("TODO")
}
