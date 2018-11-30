package core

import (
	"fmt"
	"log"
	"sync"
)

type User struct {
	lock sync.Mutex
	name string
	age  int
}

func NewMake01() {
	// new
	var i *int
	i = new(int)
	*i = 100
	log.Printf("i=%v *i=%v\n", i, *i)

	u := new(User) // or u := User{}
	u.lock.Lock()
	u.name = "Simeon"
	u.lock.Unlock()
	log.Printf("u=%v\n", u)

	// make
	dic := make(map[string]string)
	dic["A"] = "abcd"
	dic["B"] = "efg"
	fmt.Println(dic)
}
