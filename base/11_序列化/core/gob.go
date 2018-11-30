package core

// go序列化

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type P struct {
	X, Y, Z int
	Name    string
}

func Gob() {
	var network bytes.Buffer
	fmt.Printf("network: %v %T len=%v\n", network, network, network.Len())
	enc := gob.NewEncoder(&network)
	dev := gob.NewDecoder(&network)

	p1 := P{123, 456, 789, "p1"}
	err := enc.Encode(p1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("after encode: %v\n", network)
	}

	var p2 P
	err = dev.Decode(&p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("after decode: %v\n", p2)
	}
}
