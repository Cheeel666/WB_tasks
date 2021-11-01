package utils

import "sync"

// AddBlocked users to set
func AddBlocked(usersID []int) sync.Map {
	var res sync.Map
	for _, v := range usersID {
		res.Store(v, true)
	}
	return res
}
