package core

import (
	"fmt"
)

func counter() func() int {
	count := 0
	return func() int {
		count++
		fmt.Println(count)
		return count
	}
}

// Fun02 函数的闭包
// Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
func Fun02() {
	myCounter := counter()

	for i := 0; i < 10; i++ {
		myCounter()
	}
}
