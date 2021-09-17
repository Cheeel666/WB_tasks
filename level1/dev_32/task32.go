package task32

import (
	"fmt"
	"time"
)

func sleep(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}

func Run() {
	fmt.Println("started sleeping")
	sleep(1)
	fmt.Println("Finished sleeping")
}
