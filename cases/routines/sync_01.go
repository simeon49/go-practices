package routines

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex
// 我们已经看到信道非常适合在各个 Go 程间进行通信。

// 但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？

// 这里涉及的概念叫做 互斥（mutual exclusion） ，我们通常使用 互斥锁（Mutex） 这一数据结构来提供这种机制。

// Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：

// Lock
// Unlock

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) int {
	c.mux.Lock()
	c.v[key]++
	defer c.mux.Unlock()
	return c.v[key]
}

func (c *SafeCounter) Value(key string) int {
	return c.v[key]
}

func Sync01() {
	start := time.Now().Unix()
	key := "somekey"
	counter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.Inc(key)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(counter.Value(key))
	end := time.Now().Unix()
	fmt.Printf("time use: %v\n", end-start)
}
