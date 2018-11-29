package structures

import (
	"log"
)

type base struct {
	name string
}

type myStruct struct {
	*base
}

func Struct02() {
	m1 := myStruct{base: &base{name: "abc"}}
	log.Printf("m1=%v\n", m1)
	log.Printf("name=%s\n", m1.name)
	log.Printf("name=%s\n", m1.base.name)
}
