package utils

// AddBanned users to set
func AddBanned(usersID []int) map[int]bool {
	res := make(map[int]bool)
	for _, v := range usersID {
		res[v] = true
	}
	return res
}
