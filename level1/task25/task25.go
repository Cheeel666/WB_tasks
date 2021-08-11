package task25

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func (counter *Counter) inc(index int) {
	counter.mutex.Lock()
	counter.count++
	counter.mutex.Unlock()
	fmt.Println(index, counter.count)
}

func Run() {
	counter := Counter{count: 0}
	go counter.inc(0)
	go counter.inc(1)
	go counter.inc(2)
	go counter.inc(3)
	go counter.inc(4)
	time.Sleep(time.Second)
}
