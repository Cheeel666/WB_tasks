package theory

func Task4() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// чтение из открытого и пустого канала приведет к блокировке goрутины,
	for n := range ch {
		println(n)
	}
}
