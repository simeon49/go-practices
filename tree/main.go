package main

import (
	"fmt"

	"./core"
)

func doWalk(t *core.Tree, ch chan int) {
	if t.Left != nil {
		doWalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		doWalk(t.Right, ch)
	}
}

func Walk(t *core.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}
	doWalk(t, ch)
	close(ch)
}

func Same(t1, t2 *core.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if ok1 == false {
			return true
		}
	}
}

func main() {
	// fmt.Println(core.New(1))
	// fmt.Println(core.New(2))

	t1 := core.New(1)
	t2 := core.New(2)
	fmt.Printf("%v\n%v\n  %v\n\n", t1, t2, Same(t1, t2))
}
