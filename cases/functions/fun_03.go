package functions

import (
	"fmt"
)

func fibonacci() func() int {
	a, b := 0, 1
	index := 0
	return func() int {
		defer func() {
			index++
		}()
		if index == 0 {
			return 0
		} else if index == 1 {
			return 1
		}
		c := a + b
		a, b = b, c
		return c
	}
}

func Fun03() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
