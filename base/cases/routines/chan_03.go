package routines

// range 和 close
// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

// v, ok := <-ch
// 之后 ok 会被设置为 false。

// 循环 for i := range c 会不断从信道接收值，直到它被关闭。
// 注意： 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
// 还要注意： 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有值需要发送的时候才有必要关闭，例如终止一个 range 循环。

import "fmt"

func fibonacci(count int, c chan int) {
	x, y := 0, 1
	for i := 0; i < count; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Chan03() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}