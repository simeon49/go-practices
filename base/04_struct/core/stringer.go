package core

import (
	"fmt"
)

type IPAddr struct {
	p0 uint8
	p1 uint8
	p2 uint8
	p3 uint8
}

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr.p0, ipAddr.p1, ipAddr.p2, ipAddr.p3)
}

func Stringer() {
	ipAddr := IPAddr{1, 2, 3, 4}
	fmt.Println(ipAddr)
}
