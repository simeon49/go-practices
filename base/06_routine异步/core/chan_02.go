package core

import (
	"fmt"
)

// 带缓冲的信道
// 信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
// ch := make(chan int, 100)
// 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
func Chan02() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
