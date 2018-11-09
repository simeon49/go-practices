package routines

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Routine01() {
	go say("first")
	say("second")
	go say("third")
	time.Sleep(3 * time.Second)
}
