package utils

// AddBlocked users to set
func AddBlocked(usersID []int) map[int]bool {
	res := make(map[int]bool)
	for _, v := range usersID {
		res[v] = true
	}
	return res
}
