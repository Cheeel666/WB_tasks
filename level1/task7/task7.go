package task7

import (
	"fmt"
	"sync"
	"time"
)

func fillMap(mp map[int]int, i int, mut *sync.Mutex) {
	mut.Lock()
	mp[i] = i
	mut.Unlock()
}

func Run() {
	mp := map[int]int{}
	mut := sync.Mutex{}

	for i := 0; i < 10; i++ {
		go fillMap(mp, i, &mut)
	}

	time.Sleep(time.Second)
	mut.Lock()
	fmt.Println(mp)
	mut.Unlock()
}
