package task13

import (
	"fmt"
	"sync"
)

func Run() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// go func(wg *sync.WaitGroup, i int) {
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i) // Программа упадет. Надо передавать ссылку
		// }(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
