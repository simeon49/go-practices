package core

import (
	"fmt"
)

// select 语句
// select 语句使一个 Go 程可以等待多个通信操作。

// select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

func fibonacci_04(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Chan04() {
	c := make(chan int)
	quit := make(chan bool)

	go fibonacci_04(c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- true
}
