package core

import (
	"fmt"
)

// 信道
// 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

func sum(s []int, ch chan int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	ch <- sum
}

func Chan01() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// 信道
	ch := make(chan int)
	go sum(s[len(s)/2:], ch)
	go sum(s[:len(s)/2], ch)
	x, y := <-ch, <-ch
	fmt.Printf("%v %v sum: %v\n", x, y, x+y)
}
